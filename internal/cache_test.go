package pokecache

import (
	"testing"
	"time"
)

func TestGet(t *testing.T) {
    cases := []struct {
        name     string
        key      string
        expected bool
    }{
        {name: "missing key", key: "nope", expected: false},
    }

    cache := NewCache(5 * time.Second)
    for _, c := range cases {
        _, ok := cache.Get(c.key)
        if ok != c.expected {
            t.Errorf("%s: expected ok=%v, got ok=%v", c.name, c.expected, ok)
        }
    }
}
