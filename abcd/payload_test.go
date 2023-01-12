package abcd

import (
	"github.com/brian20373/gogogrpc/gen/pb/gogogrpc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPayload(t *testing.T) {
	got := FixedPayload()
	want := &gogogrpc.HelloRequest{
		Name: "Brian",
	}

	assert.Equal(t, want, got)
}
