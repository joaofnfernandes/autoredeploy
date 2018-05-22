package main

import (
	"bytes"
	"log"
	"net/http"
	"os"

	"github.com/joaofnfernandes/autoredeploy/pkg/webhook"
	"github.com/urfave/cli"
)

var apiServerCfg ApiServerConfig

func main() {
	app := cli.NewApp()
	app.Name = "API server"
	app.Flags = flags
	app.Action = serve
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("[API] error setting up server: %s", err)
	}
}

func serve(c *cli.Context) error {
	apiServerCfg = ApiServerConfigFromContext(c)
	log.Printf("[API] listening on %s\n", apiServerCfg.ServerConfig.String())

	http.HandleFunc("/", addToMessageQueue)
	err := http.ListenAndServe(apiServerCfg.ServerConfig.String(), nil)
	if err != nil {
		return err
	}
	return nil
}

func addToMessageQueue(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()

	webhook := webhook.Unmarshal(body)
	if webhook.IsValid() {
		err := Write(apiServerCfg, body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		} else {
			log.Printf("[API] wrote to message queue: %s\n", body)
			w.WriteHeader(http.StatusCreated)
		}
	} else {
		log.Printf("[API] Invalid webhook: %s\n", body)
		w.WriteHeader(http.StatusBadRequest)
	}

}
