package main

import (
	"TP-ETL/server"
	"TP-ETL/util/common"
	"TP-ETL/util/config"
	"TP-ETL/util/kafka"
	"TP-ETL/util/mongo"
	"fmt"
	"log"
	"path"
	"path/filepath"
	"runtime"

	k "github.com/confluentinc/confluent-kafka-go/kafka"
)

func rootPath() (rootPath string) {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	rootPath = filepath.Dir(d)
	return
}

func main() {
	defer func() {
		log.Printf("Start graceful shutdown")
	}()

	// init root path
	p := rootPath()
	log.Println(p)

	// loading env variables
	c := config.Init(p)

	// mongo db
	_, insert := mongo.New(
		c.Mongo.Host,
		c.Mongo.Port,
		c.Mongo.DB,
		c.Mongo.User,
		c.Mongo.Password,
		c.Mongo.Collection,
	)

	// new kafka producer instance
	// w := kafka.NewProducer(c.Kafka.Brokers, c.Kafka.Topic)
	// defer w.Close()

	// new kafka consumer instance
	consumer := kafka.NewConsumer(c.Kafka.Brokers, c.Kafka.GroupID)
	defer consumer.Close()

	// consuming data to mongo DB
	log.Printf("Start subscribe TP Kafka, and export to mongo DB")
	common.Async(func() {
		kafka.StartConsume(
			c.Kafka.Pollms,
			consumer,
			[]string{
				c.Kafka.Topic,
			},
			func(msg *k.Message) {
				var data []interface{}
				data = append(data, msg)
				insert(data)
			},
		)
	})

	// new http server
	s := server.NewHttp()
	s.Run(fmt.Sprintf(":%s", c.Server.Http.Port))
}
