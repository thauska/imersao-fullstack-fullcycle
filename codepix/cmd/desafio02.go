/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/cobra"
	"github.com/thauska/imersao-fullstack-fullcycle/codepix/consumer"
	"github.com/thauska/imersao-fullstack-fullcycle/codepix/producer"
)

// desafio02Cmd represents the desafio02 command
var desafio02Cmd = &cobra.Command{
	Use:   "desafio02",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		deliveriChan := make(chan kafka.Event)
		go producer.DeliveryWatch(deliveriChan)

		kafkaProducer := producer.NewKafkaProducer()
		go sendInfinityMsgs(kafkaProducer, deliveriChan)

		consumer.ConsumeMsgs(kafkaProducer, deliveriChan)
	},
}

func sendInfinityMsgs(kafkaProducer *kafka.Producer, deliveryChan chan kafka.Event) {
	for {
		time.Sleep(1 * time.Second)
		producer.PublishMsg("Echo!", "desafio02", kafkaProducer, deliveryChan)
	}
}

func init() {
	rootCmd.AddCommand(desafio02Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// desafio02Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// desafio02Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
