package main

import (
	"context"
	"getresponse/internal/handlers"
	"getresponse/internal/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"

	kafkago "github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

func main() {
	// preparing config file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("cannot read from a config")
	}

	// Logs
	level, ok := viper.Get("logger.loglevel").(string)
	if !ok {
        level = "debug"
    }
	loglevel, err := log.ParseLevel(level)
    if err != nil {
        loglevel = log.DebugLevel
    }
    log.SetLevel(loglevel)
	log.SetLevel(loglevel)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)

	// DB
	db, err := repository.NewDB()
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("cannot ping db: %v", err)
	}

	// Kafka
	ctx := context.Background()
	messages := make(chan kafkago.Message, 1000)
	messageCommit := make(chan kafkago.Message, 1000)
	g, ctx := errgroup.WithContext(ctx)
	writer := repository.NewKafkaWriter(messages, messageCommit)
	g.Go(func() error { return writer.WriteMessages(ctx) })

	dao := repository.NewDAO(db)
	webhook_handler := handlers.NewWebhook(dao, messages)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/v1/webhook", webhook_handler.Webhook())

	http.ListenAndServe(":3000", r)
}
