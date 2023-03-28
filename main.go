package main

import (
	"github.com/biosvos/resource-checker-go/infra/kubernetes"
	"github.com/biosvos/resource-checker-go/infra/ui"
	"log"
)

func main() {
	monitor, err := kubernetes.NewClient()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	app := ui.NewCli(monitor)
	app.Run()
}
