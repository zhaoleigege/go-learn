package algorithm

import (
	"github.com/google/uuid"
	"testing"
)

func TestUUID(t *testing.T) {
	uid := uuid.New()
	t.Log(uid)
}
