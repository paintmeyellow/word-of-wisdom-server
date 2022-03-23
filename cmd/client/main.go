package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
	"word-of-wisdom/internal/entity"
	"word-of-wisdom/pkg/tcpclient"
)

func main() {
	conn, err := tcpclient.Connect(os.Getenv("SERVER_ADDR"))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	for {
		hc, err := requestChallenge(conn)
		if err != nil {
			log.Println(err)
			continue
		}

		nonce, _ := hc.Compute()
		hc.Nonce = nonce

		quote, err := sendForVerification(conn, hc)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(quote)
		fmt.Println()

		<-time.Tick(2 * time.Second)
	}
}

func requestChallenge(conn *tcpclient.Conn) (*entity.Hashcash, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	msg, err := conn.RequestContext(ctx, "request_challenge", nil)
	if err != nil {
		return nil, err
	}
	var hc entity.Hashcash
	if err = json.Unmarshal(msg.Data, &hc); err != nil {
		return nil, err
	}
	return &hc, nil
}

func sendForVerification(conn *tcpclient.Conn, hc *entity.Hashcash) (string, error) {
	data, err := json.Marshal(hc)
	if err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	msg, err := conn.RequestContext(ctx, "verify_challenge", data)
	return string(msg.Data), nil
}
