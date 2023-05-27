package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"api-server/internal/infrastructures/application"
	"api-server/internal/infrastructures/server"
	"api-server/internal/router"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Panicln("failed to load env vars")
	}

	app, err := application.GetApplication()
	if err != nil {
		log.Panicln(err.Error())
	}

	srvConnStr := fmt.Sprintf(":%s", app.GetConfig().GetApiPort())
	srv := server.GetServer().WithAddr(srvConnStr).WithRouter(router.GetMuxRouter(app).InitRouter())

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			log.Panicln("HTTP server listen: %s\n", err)
		}
	}()

	sig := <-c

	log.Println("Caught SIGTERM %b, gracefully shutting down resources\n", sig)

	ctxx, _ := context.WithTimeout(context.Background(), 3*time.Minute)

	if err := srv.Close(ctxx); err != nil {
		log.Panicln("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")

}
