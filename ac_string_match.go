package acstringmatch

type ACStringMatch struct {
	tree *acTree
}

func NewAC(banWord []string) *ACStringMatch {
	ac := &ACStringMatch{}

	ac.tree = newACTree(banWord)

	ac.tree.buildACTree()

	return ac
}

func (a *ACStringMatch) Match(word string) int {
	return a.tree.query(word)
}

func (a *ACStringMatch) Insert(word string) {
	a.tree.insert(word)
	a.tree.buildACTree()
}

func (a *ACStringMatch) PrintTree() {
}
