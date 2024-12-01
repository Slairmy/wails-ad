package archery

import "context"

type Archery struct {
	ctx context.Context
}

func NewArchery() *Archery {
	return &Archery{}
}
