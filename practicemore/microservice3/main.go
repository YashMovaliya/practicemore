package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/YashMovaliya/practicemore/microservice2/handlers"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {
	env.Parse()
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	s := &http.Server{
		Addr:         *bindAddress,
		Handler:      sm,
		ErrorLog:     l,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		l.Println("Starting Server on port 9090")
		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error Starting Server: %s\n", err)
			os.Exit(1)
		}
	}()

	/*sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)*/
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
}
