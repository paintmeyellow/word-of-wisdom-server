package app

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"word-of-wisdom/config"
	"word-of-wisdom/internal/controller/tcp"
	"word-of-wisdom/internal/usecase"
	"word-of-wisdom/internal/usecase/repo"
	"word-of-wisdom/pkg/logger"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	r := repo.NewHashcashRepo()

	ln, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - net.Listen: %w", err))
	}

	challengeUseCase := usecase.ChallengeUseCase{Repo: r}

	h := tcp.NewHandler(l, &challengeUseCase)

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				l.Error(err)
				continue
			}
			l.Info("connected", conn.RemoteAddr())
			h.HandleConnection(conn)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	}
}
