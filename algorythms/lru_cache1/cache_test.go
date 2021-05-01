package lru_cache1

import (
	"strings"
	"testing"
)

func TestCache(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		capacity   int
		operations []string
	}{
		{
			name:     "one item",
			capacity: 1,
			operations: []string{
				"set 1 1",
				"set 2 2",
				"get 1 -1",
				"set 2 2",
			},
		},
		{
			name:     "simple check",
			capacity: 2,
			operations: []string{
				"set 1 1",
				"set 2 2",
				"get 1 1",
				"set 3 3",
				"get 2 -1",
				"set 4 4",
				"get 1 -1",
				"get 3 3",
				"get 4 4",
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			cache, _ := NewLRUCache(tc.capacity)

			for i, op := range tc.operations {
				t.Logf("step %d: %v", i, cache.items)

				step := strings.Split(op, " ")
				if step[0] == "set" {
					cache.Set(step[1], step[2])
				} else if step[0] == "get" {
					item, ok := cache.Get(step[1])

					if ok && step[2] == "-1" {
						t.Errorf("step %d: expected no item, got %s", i, item.(string))

						return
					}

					if !ok && step[2] != "-1" {
						t.Errorf("step %d: expected %s got nothing", i, step[2])

						return
					}

					if !ok && step[2] == "-1" {
						return
					}

					if item.(string) != step[2] {
						t.Errorf("step %d: expected %s got %s", i, step[2], item.(string))
					}
				} else {
					t.Fatalf("unexpected test command %s", step[0])
				}
			}
		})
	}
}
