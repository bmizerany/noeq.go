package noeq

import (
	"github.com/bmizerany/assert"
	"os"
	"testing"
)

func TestClientGen(t *testing.T) {
	c, err := New(os.Getenv("NOEQ_TOKEN"), "localhost:4444")
	if err != nil {
		t.Fatal(err)
	}

	const n = 10
	ids, err := c.Gen(n)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < n; i++ {
		assert.NotEqual(t, uint64(0), ids[0], i)
	}
}
