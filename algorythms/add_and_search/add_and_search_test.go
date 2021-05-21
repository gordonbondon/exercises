package add_and_search

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestCache(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		operations []string
	}{
		{
			name: "simple test",
			operations: []string{
				"add bad",
				"add dad",
				"add mad",
				"search pad false",
				"search bad true",
				"search .ad true",
				"search b.. true",
			},
		},
		{
			name: "interchange commands",
			operations: []string{
				"add at",
				"add and",
				"add an",
				"add add",
				"search a false",
				"search .at false",
				"add bat",
				"search .at true",
				"search an. true",
				"search a.d. false",
				"search b. false",
				"search a.d true",
				"search . false",
			},
		},
		{
			name: "more commands",
			operations: []string{
				"add ran",
				"add rune",
				"add runner",
				"add runs",
				"add add",
				"add adds",
				"add adder",
				"add addee",
				"search r.n true",
				"search ru.n.e false",
				"search add true",
				"search add. true",
				"search adde. true",
				"search .an. false",
				"search ...s true",
				"search ....e. true",
				"search ....... false",
				"search ..n.r false",
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			dict := NewDictionary()

			for i, op := range tc.operations {
				t.Logf("step %d: %s", i, op)

				command := strings.Split(op, " ")
				if len(command) > 3 || len(command) < 2 {
					t.Fatalf("wrong test command: %s", op)
				}

				switch command[0] {
				case "add":
					dict.Add(command[1])
					printTree(dict.root, 0)

				case "search":
					if len(command) < 3 {
						t.Fatalf("wrong test command: %s", op)
					}

					expected, err := strconv.ParseBool(command[2])
					if err != nil {
						t.Fatalf("wrong test command: %s", op)
					}

					actual := dict.Search(command[1])

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

func printTree(root *node, step int) {
	indent := ""

	if step == 0 {
		fmt.Println("tree:")
	}

	for i := 0; i < step; i++ {
		indent = indent + "  "
	}

	for c, n := range root.next {
		fmt.Println(indent + "\"" + c + "\"")
		printTree(n, step+1)
	}

	if step == 0 {
		fmt.Println("")
	}
}
