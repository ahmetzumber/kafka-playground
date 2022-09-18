package main

import (
	"context"
	"fmt"
	"github.com/ahmetzumber/kafka-playground/utils"
	"github.com/segmentio/kafka-go"
	"time"
)

func main() {
	conn, _ := kafka.DialLeader(context.Background(), utils.NETWORK, utils.ADDRESS, utils.TOPIC_NAME, utils.PARTITION)
	conn.SetReadDeadline(time.Now().Add(time.Second * 10))

	messages := conn.ReadBatch(1e3, 1e9)
	bytes := make([]byte, 1e6)
	for {
		_, err := messages.Read(bytes)
		if err != nil {
			break
		}
		fmt.Println(string(bytes))
	}
}

