package main


import (
	"encoding/json"
	"fmt"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)


func Consume() {

	var emp employee
	con,_ := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "testGroup",
		"auto.offset.reset": "earliest",
	})

	con.Subscribe(topic,nil)
	for {
		msg, err := con.ReadMessage(-1)
		if err == nil {
			if err:=json.Unmarshal(msg.Value,&emp);err!=nil {
				panic(err)
			}
			fmt.Printf("%+v",emp)
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}

		

	}
	con.Close()
	
}