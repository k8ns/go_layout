package main

import (
	"article/internal/article"
	"article/internal/endpoints"
	"article/internal/httproutes"
	"article/internal/myslq"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	e := &endpoints.AddArticleLog{
		Endpoint: &endpoints.AddArticleBusiness{
			Service: &article.Service{
				Storage: &myslq.Storage{},
			},
		},
	}

	et := &endpoints.AddArticleTime{
		Endpoint: e,
	}

	mux := httproutes.GetRoutes(et)
	mux = httproutes.LoggingMiddleware(mux)

	ctx, interrupt := context.WithCancel(context.Background())

	gracefulListenAndServe(ctx, ":80", mux)

	quit := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	interrupt()

    time.Sleep(3 * time.Second)
    log.Println("app shutdown")
}

func gracefulListenAndServe(ctx context.Context, addr string, handler http.Handler) {
	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				if err := server.Shutdown(ctx); err != nil {
					log.Fatal("Server Shutdown error:", err)
				}
				log.Println("Server switched off")
				return
			}
		}
	}()
}
