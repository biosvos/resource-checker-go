package ui

import (
	"github.com/biosvos/resource-checker-go/flow/monitor"
	"log"
)

type Cli struct {
	client monitor.Client
}

func NewCli(client monitor.Client) *Cli {
	return &Cli{client: client}
}

func (c *Cli) Run() {
	resources, err := c.client.List("", "v1", "pod", "")
	if err != nil {
		log.Fatalf("%+v", err)
	}

	for _, resource := range resources {
		log.Println(resource.Namespace, resource.Kind, resource.Name)
	}
}
