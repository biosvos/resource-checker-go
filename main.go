package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/biosvos/resource-checker-go/flow/flower"
	"github.com/biosvos/resource-checker-go/infra/kubernetes"
	"github.com/biosvos/resource-checker-go/infra/unstructure"
	"log"
)

func prettyPrint(b []byte) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Print(out.String())
}

func main() {
	// monitor := memory.NewMemory()
	// monitor.AddResources(chocoFailedDeployJson, chocoFailedReplicasetJson, chocoFailedPodManifest)
	monitor, err := kubernetes.NewClient()
	if err != nil {
		panic(err)
	}
	flow := flower.NewFlow(monitor, unstructure.NewFactory())
	workload, err := flow.GetFamily(&flower.Resource{
		GroupVersionKind: flower.GroupVersionKind{
			Group:   "apps",
			Version: "v1",
			Kind:    "Deployment",
		},
		Namespace: "wow",
		Name:      "deployment-not-owner",
	})
	if err != nil {
		panic(err)
	}
	for _, resource := range workload {
		log.Println(resource.Group, resource.Version, resource.Kind, resource.Namespace, resource.Name)
	}
}
