package acstringmatch

type node struct {
	word      rune
	child     map[rune]*node
	failIndex *node
	isTail    bool
}

type acTree struct {
	root     *node
	treeSize int
}

func newACTree(banWord []string) *acTree {
	t := &acTree{
		root: &node{
			child: map[rune]*node{},
		},
	}

	for _, b := range banWord {
		t.insert(b)
	}

	return t
}

func (t *acTree) insert(word string) {
	cur := t.root

	for _, w := range word {
		n, ok := cur.child[w]

		if !ok {
			n = &node{
				word:  w,
				child: map[rune]*node{},
			}
			cur.child[w] = n
			t.treeSize++
		}

		cur = n
	}
}

func (t *acTree) query(word string) int {

	cur := t.root

	var count int

	for _, w := range word {
	fail:
		n, ok := cur.child[w]
		if !ok {
			if cur.failIndex == t.root {
				cur = t.root
				continue
			} else {
				cur = cur.failIndex
				goto fail
			}
		} else {

			if n.isTail {
				count++
			}

			cur = n
		}
	}

	return count
}

func (t *acTree) buildACTree() {
	queue := make([]*node, 0, t.treeSize)
	queue = append(queue, t.root)

	parentNodeMap := make(map[rune]*node)

	for len(queue) > 0 {
		cur := queue[0]

		for _, c := range cur.child {
			queue = append(queue, c)
		}

		if len(cur.child) == 0 {
			cur.isTail = true
		}

		queue = queue[1:]

		if cur == t.root {
			cur.failIndex = t.root
		} else {
			parent, exist := parentNodeMap[cur.word]

			if !exist {
				cur.failIndex = t.root
				parentNodeMap[cur.word] = cur
			} else {
				cur.failIndex = parent
			}
		}
	}
}
