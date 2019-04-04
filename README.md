## NATS-STREAMING-DEMO


#### Prerequisites

- `go get -u github.com/masudur-rahman/nats-streaming-demo`

#### Run nats-streaming-server in one window

- `cd /home/$USER/go/src/github.com/masudur-rahman/nats-streaming-demo`
- `./nats-streaming-demo`

#### Subscribe to NATS-Streaming in second window

- `cd /home/$USER/go/src/github.com/masudur-rahman/nats-streaming-demo`
- `go run subscribe/main.go <subject>` | example: `go run subscribe/main.go nats-streaming`

#### Publish message using NATS-Streaming in third window

- `cd /home/$USER/go/src/github.com/masudur-rahman/nats-streaming-demo`
- `go run publish/main.go <subject> <msg>` | example: `go run subscribe/main.go nats-streaming "Welcome to nats streaming"`


###### Now you are all set. You will find every unread message in `second window` regarding the provided `Subject`.
  