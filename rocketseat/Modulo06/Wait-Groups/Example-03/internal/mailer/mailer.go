package mailer

import (
	"log"
	"sync"
	"time"

	"github.com/go-mail/mail/v2"
)

type Sender struct {
	dialer *mail.Dialer
	addr   string
}

func NewSender(host string, port int, username, password, addr string) Sender {
	return Sender{
		dialer: mail.NewDialer(host, port, username, password),
		addr:   addr,
	}
}

func (s Sender) Send(recipient string, wg *sync.WaitGroup) error {
	defer wg.Done()
	msg := mail.NewMessage()
	msg.SetHeader("To", recipient)
	msg.SetHeader("From", s.addr)
	msg.SetHeader("Subject", "An email from a Awesome SMTP wrapper")
	msg.SetHeader("text/plain", "Hey this is a test email from our microservice! I hope tou get this!")
	log.Printf("started sending")
	time.Sleep(5 * time.Second)
	defer log.Printf("Finish sending")
	return s.dialer.DialAndSend(msg)
}
