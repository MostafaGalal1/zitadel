package gomap

import (
	"context"
	"testing"
	"time"

	"github.com/zitadel/zitadel/internal/cache"
)

type testIndex int

const (
	testIndexID testIndex = iota
	testIndexName
)

var testIndices = []testIndex{
	testIndexID,
	testIndexName,
}

type testObject struct {
	id    string
	names []string
}

func (o *testObject) Keys(index testIndex) []string {
	switch index {
	case testIndexID:
		return []string{o.id}
	case testIndexName:
		return o.names
	default:
		return nil
	}
}

func (o *testObject) Value() *testObject {
	return o
}

func Test(t *testing.T) {
	c := NewCache[testIndex, string, *testObject](context.Background(), testIndices, cache.CacheConfig{
		MaxAge:     time.Second,
		LastUseAge: time.Second / 4,
	})
	defer c.Close(context.Background())
	c.Set(context.Background(), &testObject{})
}
