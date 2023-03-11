package main

import (
	"log"

	"github.com/spf13/cobra"
)

func main() {
	root := cobra.Command{
		Use:   "ktl",
		Short: "A command line tool to interact with Kafka",
		Long: `A Kafka command line interface
		
		ktl helps to enable Kafka administration, message producing and message consuming.
		The client uses franz-go library to interact with Kafka APIs
		`,
	}

	root.AddCommand()
	if err := root.Execute(); err != nil {
		log.Fatalln(err)
	}
}
