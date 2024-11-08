package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var myUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/health", healthHandler)
	router.HandleFunc("/shell", shellHandler)

	server := &http.Server{
		Addr:    ":5000",
		Handler: router,
	}

	log.Println("INFO: server started at port 5000")
	err := server.ListenAndServe()
	if err != nil {
		log.Println("ERROR: starting the server")
	}
}

// Health Check handler
func healthHandler(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "API Health is OK.")
}

func shellHandler(w http.ResponseWriter, request *http.Request) {
	connection, err := myUpgrader.Upgrade(w, request, nil)
	if err != nil {
		log.Println("ERROR: Unable to upgrade the connection to Web Socket", err)
		return
	}

	for {
		mt, messageByte, err := connection.ReadMessage()
		if err != nil {
			log.Println("ERROR: Unable to read the message", err)
			return
		}

		// execute the shell commands
		command := exec.Command("/bin/sh", "-c", string(messageByte))
		output, err := command.CombinedOutput()
		if err != nil {
			log.Println("ERROR: unable to execute the command", err, command)
		}

		err = connection.WriteMessage(mt, output)
		if err != nil {
			log.Println("WRITE:", err)
			break
		}

	}

}
