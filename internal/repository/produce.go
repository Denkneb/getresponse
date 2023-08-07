package repository

import (
	"context"
	"fmt"
	"github.com/spf13/viper"

	kafkago "github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

type Producer interface {
}

type producer struct {
	Writer        *kafkago.Writer
	Messages      chan kafkago.Message
	messageCommit chan kafkago.Message
}

func NewKafkaWriter(messages, messageCommit chan kafkago.Message) *producer {
	host := viper.Get("kafka.host").(string)
	port := viper.Get("kafka.port").(int)
	topic := viper.Get("kafka.producer.topic").(string)
	writer := &kafkago.Writer{
		Addr:  kafkago.TCP(fmt.Sprintf("%s:%d", host, port)),
		Topic: topic,
	}
	return &producer{
		Writer:        writer,
		Messages:      messages,
		messageCommit: messageCommit,
	}
}

func (k *producer) WriteMessages(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case m := <-k.Messages:
			err := k.Writer.WriteMessages(ctx, m)
			if err != nil {
				log.Warning(err)
			}
			log.Debug("sent message")

			select {
			case <-ctx.Done():
			case k.messageCommit <- m:
			}
		}
	}
}
