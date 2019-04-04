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
		api.PubClientID,
		stan.NatsURL(stan.DefaultNatsURL),
		stan.ConnectWait(10*time.Second),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalln("Connection lost, reason:", reason)
		}),
	)
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, stan.DefaultNatsURL)
	}
	defer logCloser(conn)

	log.Printf("Connected to %s clusterID: [%s] clientID: [%s]\n", stan.DefaultNatsURL, api.ClusterID, api.PubClientID)

	args := os.Args

	subj, msg := args[1], args[2]
	log.Println(subj, msg)

	if err = conn.Publish(subj, []byte(msg)); err != nil {
		log.Fatalln("Error during publish:", err)
	}
}
