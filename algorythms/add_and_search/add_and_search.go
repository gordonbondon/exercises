package add_and_search

const (
	dotMatch = "."
	end      = ""
)

type node struct {
	next map[string]*node
}

type Dictionary struct {
	root *node
}

func NewDictionary() *Dictionary {
	dict := &Dictionary{}

	nodes := make(map[string]*node)
	dict.root = &node{next: nodes}

	return dict
}

func (d *Dictionary) Add(word string) {
	position := d.root

	for _, c := range word {
		if next, ok := position.next[string(c)]; ok {
			position = next

			continue
		} else {
			nodes := make(map[string]*node)

			next := &node{next: nodes}

			position.next[string(c)] = next
			position = next

			continue
		}
	}

	position.next[end] = &node{}
}

func (d *Dictionary) Search(word string) bool {
	position := 0
	queue := []*node{d.root}

	for len(queue) > 0 && position < len(word)+1 {
		// check if we can end the word
		if position == len(word) {
			for _, q := range queue {
				for s := range q.next {
					if end == s {
						return true
					}
				}
			}

			position++

			continue
		}

		character := string([]rune(word)[position])

		oldLen := len(queue)

		for _, q := range queue {
			for s, next := range q.next {
				if character == s || character == dotMatch {
					queue = append(queue, next)
				}
			}
		}

		queue = queue[oldLen:]

		position++
	}

	return false
}
