package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
)

var topic = "my-topic"
var reader *kafka.Reader

func main() {
	fmt.Println("Hello, World!")
	ctx := context.Background()
	go listenSignal()
	// readKafka(ctx)
	WriteKafka(ctx, "Hello, Kafka!")
}

func logf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func WriteKafka(ctx context.Context, message string) {
	fmt.Println("WriteKafka")
	writer := &kafka.Writer{
		Addr:        kafka.TCP("localhost:9092"),
		Topic:       topic,
		Balancer:    &kafka.Hash{},
		Logger:      kafka.LoggerFunc(logf),
		ErrorLogger: kafka.LoggerFunc(logf),
	}
	defer writer.Close()
	for i := 0; i < 3; i++ {
		if err := writer.WriteMessages(ctx, kafka.Message{
			Key:   []byte("1111"),
			Value: []byte(message),
		}); err != nil {
			// 第一次写操作失败，重试
			if err == kafka.LeaderNotAvailable {
				time.Sleep(time.Second * 2)
				continue
			} else {
				fmt.Printf("write message error: %v", err)
				return
			}
		}

		fmt.Println("write message success")
	}

}

func readKafka(ctx context.Context) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"localhost:9092", "localhost:9093", "localhost:9094"},
		Topic:          topic,
		Partition:      0,
		CommitInterval: time.Second * 2,
		MaxBytes:       10e6, //
		StartOffset:    kafka.FirstOffset,
	})

	// defer reader.Close()

	for {
		m, err := reader.ReadMessage(ctx)
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := reader.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

func listenSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGBUS, syscall.SIGSEGV, syscall.SIGKILL, syscall.SIGINT)
	sig := <-c
	fmt.Printf("Got signal: %v, exiting...", sig)
	if reader != nil {
		reader.Close()
	}
	os.Exit(0)
}
