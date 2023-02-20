package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	// controller "github.com/tawseefnabi/urlshortner/Controller"
	repository "github.com/tawseefnabi/urlshortner/Repository"
	// service "github.com/tawseefnabi/urlshortner/Service"
	"github.com/tawseefnabi/urlshortner/shortenurl"
	"gorm.io/gorm"
)

var (
	Database = "tinyUrl.db"
)

var db *gorm.DB

// var rdb *redis.Client
var err error

func handler(w http.ResponseWriter, r *http.Request) {
	return
}
func init() {
	db, err = shortenurl.Connect(Database)
	fmt.Println("db", db, err)
}

func main() {
	// https://github.com/gorilla/mux#graceful-shutdown
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	repository := repository.NewRepository(db)
	fmt.Println("repo", repository)
	// service := service.NewService(repository)
	// controller := controller.NewController(service)
	router := mux.NewRouter()

	router.HandleFunc("/generateTinyUrl/", handler).Methods("POST")

	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}
	fmt.Println("server running at ", srv.Addr)
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}