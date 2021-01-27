package kafkaexample

import (
	"crypto/rand"
	"time"
)

// Ulid encapsulate the way to generate ulids
func Ulid() string {
	t := time.Now().UTC()
	id := ulid.MustNew(ulid.Timestamp(t), rand.Reader)

	return id.String()
}
