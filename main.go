package main

import (
	"log"
	"os"

	"github.com/joseluisq/goipcc"

	"xojoidecom/cmd"
)

func main() {
	// Example of how compile a Xojo app via Unix IPC socket communication

	// Configure and exchange data with current socket
	ipc, err := goipcc.New(cmd.UnixSocketPath)
	if err != nil {
		log.Println("unable to communicate with socket:", err)
		os.Exit(1)
	}

	// Send many requests
	pangrama := []string{
		"{\"protocol\":2}\x00",
		"{\"script\":\"Print BuildApp(16,False)\", \"tag\":\"build\"}\x00",
	}
	for _, word := range pangrama {
		_, err := ipc.Write([]byte(word + "\n"))
		if err != nil {
			log.Fatalln("unable to write to socket:", err)
			break
		}
		log.Println("client data sent:", word)
	}

	// Listen for socket responses
	ipc.Listen(func(data []byte, err error) {
		if err != nil {
			log.Fatalln("unable to get data:", err)
		}
		log.Println("client data got:", string(data))
	})
}
