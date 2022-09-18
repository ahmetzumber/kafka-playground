package main

import (
	"context"
	"github.com/ahmetzumber/kafka-playground/utils"
	"github.com/segmentio/kafka-go"
	"time"
)

func main() {
	conn, _ := kafka.DialLeader(context.Background(), utils.NETWORK, utils.ADDRESS, utils.TOPIC_NAME, utils.PARTITION)
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))

	conn.WriteMessages(kafka.Message{ Value: []byte("catch the first event !!") })
}

