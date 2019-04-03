package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/masudur-rahman/nats-streaming-demo/api"
	"gopkg.in/macaron.v1"

	stan "github.com/nats-io/go-nats-streaming"
)

func ProcessMsg(msg *stan.Msg) {
	fmt.Println(string(msg.Data))
}

func logCloser(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Println("Close error:", err)
	}
}

func main() {
	conn, err := stan.Connect(
		api.ClusterID,
		api.ClientID,
		stan.NatsURL(stan.DefaultNatsURL),
		stan.ConnectWait(10*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer logCloser(conn)

	args := os.Args

	subj := args[1]
	log.Println(subj)

	sub, err := conn.Subscribe(
		subj,
		ProcessMsg,
	)
	defer logCloser(sub)

	m := macaron.Classic()

	m.Get("/", func() string {
		return "Hello stranger...!"
	})
	m.Run()
}
