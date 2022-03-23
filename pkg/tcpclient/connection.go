package tcpclient

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
)

const SubsChanLen = 8

const _EMPTY_ = ""

var (
	ErrConnectionClosed = errors.New("tcpclient: connection closed")
	ErrMsgNotBound      = errors.New("tcpclient: message is not bound to subscription/connection")
	ErrCancelled        = errors.New("tcpclient: context cancelled")
)

type Msg struct {
	ID    string
	Reply string
	Subj  string
	Data  []byte
	Conn  *Conn
	sync.RWMutex
}

func (m *Msg) Respond(data []byte) error {
	if m == nil || m.Conn == nil {
		return ErrMsgNotBound
	}
	m.RLock()
	conn := m.Conn
	m.RUnlock()
	return conn.publish(m.Subj, m.ID, _EMPTY_, data)
}

type MsgHandler func(m *Msg)

type Conn struct {
	conn     net.Conn
	closing  bool
	closed   bool
	subs     map[string][]chan *Msg
	respMap  map[string]chan *Msg
	respRand *rand.Rand
	subsMu   sync.RWMutex
	sync.RWMutex
}

func NewConn(c net.Conn) *Conn {
	return newConn(c)
}

func Connect(addr string) (*Conn, error) {
	c, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return newConn(c), nil
}

func (c *Conn) Close() {
	c.Lock()
	c.closing = true
	if c.closed {
		c.closing = false
		return
	}
	c.Unlock()
	c.subsMu.Lock()
	for _, subs := range c.subs {
		for _, ch := range subs {
			close(ch)
		}
	}
	c.subsMu.Unlock()
	c.Lock()
	c.conn.Close()
	c.closed = true
	c.closing = false
	c.Unlock()
}

func (c *Conn) Publish(subj string, data []byte) error {
	return c.publish(subj, _EMPTY_, _EMPTY_, data)
}

func (c *Conn) Subscribe(subj string, h MsgHandler) {
	ch := make(chan *Msg)
	c.subscribe(subj, ch)
	go func() {
		for {
			m, ok := <-ch
			if !ok {
				return
			}
			h(m)
		}
	}()
}

func (c *Conn) RequestContext(ctx context.Context, subj string, data []byte) (*Msg, error) {
	id := c.randomToken()
	c.Lock()
	mch := make(chan *Msg)
	c.respMap[id] = mch
	c.Unlock()

	if err := c.publish(subj, _EMPTY_, id, data); err != nil {
		return nil, err
	}
	var (
		msg *Msg
		err error
		ok  bool
	)
	select {
	case msg, ok = <-mch:
		if !ok {
			return nil, ErrConnectionClosed
		}
	case <-ctx.Done():
		err = ErrCancelled
	}
	c.Lock()
	delete(c.respMap, id)
	c.Unlock()

	return msg, err
}

func newConn(c net.Conn) *Conn {
	conn := Conn{
		conn:     c,
		subs:     make(map[string][]chan *Msg, SubsChanLen),
		respMap:  make(map[string]chan *Msg),
		respRand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	go conn.readLoop()
	return &conn
}

func (c *Conn) readLoop() {
	defer c.Close()
	reader := bufio.NewReader(c.conn)
	for {
		b, err := reader.ReadBytes('\n')
		switch err {
		case nil:
			var msg Msg
			if err := json.NewDecoder(bytes.NewReader(b)).Decode(&msg); err != nil {
				log.Printf("err decode message\n")
			}
			msg.Lock()
			msg.Conn = c
			msg.Unlock()
			if msg.Reply != _EMPTY_ {
				c.Lock()
				ch, ok := c.respMap[msg.Reply]
				c.Unlock()
				if ok {
					go func(ch chan *Msg) {
						if !c.closing && !c.closed {
							ch <- &msg
						}
					}(ch)
				}
			} else {
				c.subsMu.Lock()
				subs := c.subs[msg.Subj]
				c.subsMu.Unlock()
				for _, ch := range subs {
					go func(ch chan *Msg) {
						if !c.closing && !c.closed {
							ch <- &msg
						}
					}(ch)
				}
			}
		case io.EOF:
			log.Println("client closed the connection by terminating the process")
			return
		default:
			log.Printf("error: %v\n", err)
			return
		}
	}
}

func (c *Conn) publish(subj, reply, id string, data []byte) error {
	c.Lock()
	defer c.Unlock()
	if c.closed {
		return ErrConnectionClosed
	}
	if id == _EMPTY_ {
		id = c.randomToken()
	}
	b, err := json.Marshal(Msg{
		ID:    id,
		Reply: reply,
		Subj:  subj,
		Data:  data,
	})
	if err != nil {
		return err
	}
	_, err = c.conn.Write(append(b, '\n'))
	return err
}

func (c *Conn) subscribe(subj string, ch chan *Msg) {
	c.Lock()
	defer c.Unlock()
	if !c.closed {
		c.subsMu.Lock()
		c.subs[subj] = append(c.subs[subj], ch)
		c.subsMu.Unlock()
	}
}
