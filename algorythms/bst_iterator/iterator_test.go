package bst_iterator

import (
	"strconv"
	"strings"
	"testing"
)

func TestCache(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		root       *TreeNode
		operations []string
	}{
		{
			name: "simple test",
			root: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 15,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 20,
					},
				},
			},
			operations: []string{
				"next 3",
				"next 7",
				"hasNext true",
				"next 9",
				"hasNext true",
				"next 15",
				"hasNext true",
				"next 20",
				"hasNext false",
			},
		},
		{
			name: "only root",
			root: &TreeNode{
				Val: 7,
			},
			operations: []string{
				"hasNext true",
				"next 7",
				"hasNext false",
			},
		},
		{
			name: "only right",
			root: &TreeNode{
				Val: 7,
				Right: &TreeNode{
					Val: 10,
					Right: &TreeNode{
						Val: 20,
						Right: &TreeNode{
							Val: 30,
						},
					},
				},
			},
			operations: []string{
				"hasNext true",
				"next 7",
				"next 10",
				"next 20",
				"next 30",
				"hasNext false",
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			iterator := NewBSTIterator(tc.root)

			for i, op := range tc.operations {
				t.Logf("step %d: %s", i, op)

				command := strings.Split(op, " ")
				if len(command) != 2 {
					t.Fatalf("wrong test command: %s", op)
				}

				switch command[0] {
				case "next":
					expected, err := strconv.ParseInt(command[1], 10, 0)
					if err != nil {
						t.Fatalf("wrong test command: %s", op)
					}

					actual := iterator.Next()

					if actual != int(expected) {
						t.Errorf("expected %v got %v", expected, actual)
					}

				case "hasNext":
					expected, err := strconv.ParseBool(command[1])
					if err != nil {
						t.Fatalf("wrong test command: %s", op)
					}

					actual := iterator.HasNext()

					if actual != expected {
						t.Errorf("expected %v got %v", expected, actual)
					}

				default:
					t.Fatalf("wrong test command: %s", op)
				}
			}
		})
	}
}
