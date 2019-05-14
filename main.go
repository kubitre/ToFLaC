package main

import (
	"log"
	"net/http"
	"tflac_cw/kernel"
	"tflac_cw/routers"

	"github.com/rs/cors"
)

func entryPointStart() {
	entry := &routers.EntryPoint{}
	entry = entry.Init()

	handler := cors.New(
		cors.Options{
			AllowedMethods: []string{"GET", "POST"},
			AllowedHeaders: []string{"Content-Type", "Accept", "Access-Control-Allow-Origin"},
			AllowedOrigins: []string{"*"},
			Debug:          true,
		}).Handler(entry.CreateEntryPoint())

	if err := http.ListenAndServe(":9995", handler); err != nil {
		log.Fatal(err)
	}
}

func testkernekProcess() {
	ker := &kernel.Kernel{}
	ker = ker.New()

	ker.Input("flot lol \n int test;")
}

func main() {
	// entryPointStart()
	testkernekProcess()
}
