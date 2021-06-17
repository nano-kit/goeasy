package impl

import (
	context "context"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/nano-kit/goeasy/internal/ierr"
	seq "github.com/nano-kit/goeasy/servers/sequence"
)

type Sequence struct {
	s *Server
}

func (s *Sequence) Next(ctx context.Context, req *seq.NextReq, res *seq.NextRes) error {
	if req.Name == "" {
		return ierr.BadRequest("sequence name is empty")
	}
	key := s.key(req.Name)
	val, err := s.db().Incr(ctx, key).Result()
	if err != nil {
		return ierr.Storage("INCR %q: %v", key, err)
	}
	res.Value = uint64(val)
	return nil
}

func (s *Sequence) Max(ctx context.Context, req *seq.MaxReq, res *seq.MaxRes) error {
	if req.Name == "" {
		return ierr.BadRequest("sequence name is empty")
	}
	key := s.key(req.Name)
	val, err := s.db().Get(ctx, key).Result()
	if err != nil {
		return ierr.Storage("GET %q: %v", key, err)
	}
	res.Value, err = strconv.ParseUint(val, 10, 64)
	if err != nil {
		return ierr.Storage("ParseUint GET %q: %v", key, err)
	}
	return nil
}

// key returns the unique sequence key, by adding a prefix to the sequence name.
// The prefix should not conflict with other server.
func (s *Sequence) key(name string) string {
	return "seq:" + name
}

func (s *Sequence) db() *redis.Client {
	return s.s.redisDB
}
