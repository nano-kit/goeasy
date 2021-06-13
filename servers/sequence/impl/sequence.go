package impl

import (
	context "context"

	seq "github.com/nano-kit/goeasy/servers/sequence"
)

type Sequence struct{}

func (s *Sequence) Next(ctx context.Context, req *seq.NextReq, res *seq.NextRes) error {
	return nil
}

func (s *Sequence) Max(ctx context.Context, req *seq.MaxReq, res *seq.MaxRes) error {
	return nil
}
