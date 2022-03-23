package tcp

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net"
	"word-of-wisdom/internal/entity"
	"word-of-wisdom/internal/usecase"
	"word-of-wisdom/pkg/logger"
	"word-of-wisdom/pkg/tcpclient"
)

var (
	ErrSendResponse = errors.New("err send response back")
)

type Handler struct {
	log       logger.Interface
	challenge *usecase.ChallengeUseCase
}

func NewHandler(l logger.Interface, c *usecase.ChallengeUseCase) *Handler {
	return &Handler{log: l, challenge: c}
}

func (h *Handler) HandleConnection(conn net.Conn) {
	c := tcpclient.NewConn(conn)
	resource := []byte(conn.RemoteAddr().String())
	c.Subscribe("request_challenge", h.requestChallenge(resource))
	c.Subscribe("verify_challenge", h.verifyChallenge(resource))
}

func (h *Handler) requestChallenge(resource []byte) tcpclient.MsgHandler {
	return func(m *tcpclient.Msg) {
		hashcashBytes, err := h.challenge.Request(resource)
		if err != nil {
			if err = m.Respond(errorResp(err)); err != nil {
				h.log.Error(ErrSendResponse)
			}
		}
		if err = m.Respond(hashcashBytes); err != nil {
			h.log.Error(ErrSendResponse)
		}
	}
}

func (h *Handler) verifyChallenge(resource []byte) tcpclient.MsgHandler {
	return func(m *tcpclient.Msg) {
		var hc entity.Hashcash
		if err := json.Unmarshal(m.Data, &hc); err != nil {
			h.log.Error(err)
			if err = m.Respond(errorResp(err)); err != nil {
				h.log.Error(ErrSendResponse)
			}
		}
		quote, err := h.challenge.Verify(&hc, resource)
		if err != nil {
			h.log.Error(err)
			if err = m.Respond(errorResp(err)); err != nil {
				h.log.Error(ErrSendResponse)
			}
		}
		if err := m.Respond([]byte(quote)); err != nil {
			h.log.Error(ErrSendResponse)
		}
	}
}

func errorResp(err error) []byte {
	return []byte(err.Error())
}
