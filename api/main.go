package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

type ServerConfig struct {
	iface string
	port  string
}

func defaultServerConfig() *ServerConfig {
	return &ServerConfig{"0.0.0.0", "8000"}
}
func (c *ServerConfig) String() string {
	return fmt.Sprintf("%s:%s", c.iface, c.port)
}

func main() {
	cfg := defaultServerConfig()
	log.Printf("[API] listening on %s\n", cfg)

	http.HandleFunc("/", addToMessageQueue)
	err := http.ListenAndServe(defaultServerConfig().String(), nil)
	if err != nil {
		log.Fatalf("[API] error setting up server: %s")
	}
}

func addToMessageQueue(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()

	err := Write(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		log.Printf("[API] wrote to message queue: %s\n", body)
		w.WriteHeader(http.StatusOK)
	}

}
