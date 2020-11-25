package cmd

import (
	"log"
	"os"

	"github.com/joseluisq/goipcc"
)

// Execute is the main entry point of the current application.
func Execute() {
	// Example of how compile a Xojo app via Unix IPC socket communication

	// Connect to socket in oder to exchange data
	ipc := goipcc.New(UnixSocketPath)
	if err := ipc.Connect(); err != nil {
		log.Println("unable to communicate with socket:", err)
		os.Exit(1)
	}

	// Xojo IDE Communication Protocol commands in oder to compile a opened Xojo project
	instructions := []string{
		"{\"protocol\":2}\x00",
		"{\"script\":\"Print BuildApp(16,False)\", \"tag\":\"build\"}\x00",
	}
	for i, req := range instructions {
		log.Println("client data sent:", req)

		var err error
		// Skip first sent since it just specifies the protocol without return
		if i == 0 {
			_, err = ipc.Write([]byte(req), nil)
		} else {
			// Next sent will return a Xojo response after the project compilation
			_, err = ipc.Write([]byte(req), func(resp []byte, err error) {
				log.Println("client data received:", string(resp))
			})
		}
		if err != nil {
			log.Fatalln("unable to write to socket:", err)
		}
	}

	ipc.Close()
}
