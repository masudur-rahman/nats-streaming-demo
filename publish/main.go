package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/masudur-rahman/nats-streaming-demo/api"

	stan "github.com/nats-io/go-nats-streaming"
)

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

	subj, msg := args[1], []byte(args[2])
	log.Println(subj, msg)

	if err = conn.Publish(subj, msg); err != nil {
		log.Fatalln("Error during publish:", err)
	}
}
