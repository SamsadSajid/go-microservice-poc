package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"net/http"
)

type Blash struct {
	Hello string `json:"hello"`
}

func getData(c *gin.Context) {
	var body Blash


	if err := c.BindJSON(&body); err != nil {
		panic(err)
	}

	log.Println(" [x] Requesting data ", body.Hello)

	conn, err := amqp.Dial("amqp://localhost")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	corrId := randomString(32)

	err = ch.Publish(
		"",          // exchange
		"rpc_queue", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          [] byte(body.Hello),
		})
	failOnError(err, "Failed to publish a message")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var res []byte

	for d := range msgs {
		if corrId == d.CorrelationId {
			//fmt.Println(d)
			res = d.Body
			fmt.Println(string(d.Body))
			//failOnError(err, "Failed to convert body to integer")
			break
		}
	}

	c.JSON(http.StatusOK, string(res))
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

