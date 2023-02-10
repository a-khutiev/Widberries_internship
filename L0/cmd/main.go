package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/a-khutiev/Widberries_internship/internal/cache"
	"github.com/a-khutiev/Widberries_internship/internal/db"
	"github.com/a-khutiev/Widberries_internship/internal/handler"
	"github.com/a-khutiev/Widberries_internship/internal/nats"
	"github.com/a-khutiev/Widberries_internship/internal/server"
	"github.com/pkg/errors"
)

const port = ":8080"

func main() {
	ctx := context.Background()

	// подключиться к бд
	database, err := db.New(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	// создать кэш
	cacheImpl, err := cache.New(ctx, database)
	if err != nil {
		log.Fatal(err)
	}

	// подключиться к nats
	natsHandler, err := nats.New(cacheImpl)
	if err != nil {
		log.Fatal(err)
	}
	defer natsHandler.Close()

	// начать обработку сообщений
	err = natsHandler.Start()
	if err != nil {
		log.Fatal(err)
	}

	// запустить HTTP сервер
	h := handler.New(database, cacheImpl)
	srv := server.New(port, h)
	go func() {
		if err = srv.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error occurred while running http server: %v", err)
		}
	}()

	log.Printf("service started on port %s", port)

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	log.Printf("service shutting down")
	if err = srv.Shutdown(ctx); err != nil {
		log.Printf("failed to shut down: %v", err)
	}
}
