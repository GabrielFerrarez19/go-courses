package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"Example-03/internal/mailer"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	var wg sync.WaitGroup
	r := chi.NewMux()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	s := mailer.NewSender("sandbox.smtp.mailtrap.io", 2525, "c1cb02087c67a2", "0e1a0a13f442d4", "myEmail@goabriel.net")

	r.Route("/api", func(r chi.Router) {
		r.Route("/mail", func(r chi.Router) {
			r.Route("/send", func(r chi.Router) {
				r.Post("/{email}", func(w http.ResponseWriter, r *http.Request) {
					userMail := chi.URLParam(r, "email")
					wg.Add(1)
					go s.Send(userMail, &wg)
				})
			})
		})
	})

	go http.ListenAndServe("localhost:6767", r)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Println("Got a signal to stop... CLEANING UP!")
	wg.Wait()
}
