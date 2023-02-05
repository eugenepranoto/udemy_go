package facade

// membuat 1 struct yang memiliki fungsi
// dimana fungsi didalamnya itu banyak
// dan interaksi dengan modul modul lain
// jadi mensimplify. kita tidak perlu tau fungsi didalamnya
// cukup panggil struct facade ini lalu panggil fungsinys

import (
	"fmt"
	"udemy/go/facade/gofood"
)

type TicketType string

const (
	GoFood TicketType = "goFood"
	GoRide TicketType = "goRide"
	GoSend TicketType = "goSend"
)

type Ticket struct {
	UserEmail string
	Type      TicketType
	Message   string
}

type Response struct {
	Status  string
	Message string
}

func (t Ticket) HandleTicket() *Response {
	var message string
	switch t.Type {
	case GoFood:
		message = gofood.CheckMessage(t.Message)
	case GoRide:
	case GoSend:
	}
	return &Response{Message: message}
}

func init() {
	ticket := Ticket{UserEmail: "eugene@gmail.com", Type: GoSend, Message: "TEST"}
	res := ticket.HandleTicket()
	fmt.Println(res)
}
