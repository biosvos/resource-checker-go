package main

import (
	"github.com/biosvos/resource-checker-go/flow/flower"
	"github.com/biosvos/resource-checker-go/infra/memory"
	"github.com/biosvos/resource-checker-go/infra/unstructure"
	"log"
)

func main() {
	monitor := memory.NewMemory()
	monitor.AddResources(chocoFailedDeployJson, chocoFailedReplicasetJson, chocoFailedPodManifest)
	flow := flower.NewFlow(monitor, unstructure.NewFactory())
	workload, err := flow.GetFamily(&flower.Resource{
		GroupVersionKind: flower.GroupVersionKind{
			Group:   "apps",
			Version: "v1",
			Kind:    "Deployment",
		},
		Namespace: "wow",
		Name:      "choco",
	})
	if err != nil {
		panic(err)
	}
	log.Println(workload)
}
