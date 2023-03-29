package unstructure

import "github.com/biosvos/resource-checker-go/flow/familiar"

var _ familiar.Familiar = &Empty{}

type Empty struct {
}

func (e *Empty) NeedMore() []*familiar.Id {
	return nil
}
