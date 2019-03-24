package queue

type Collection []string

func (c *Collection) Add(s string) {
	*c = append(*c, s)
}

func (c *Collection) Pop() string {
	var out string
	if !c.IsEmpty() {
		out = (*c)[len(*c)-1]
		*c = (*c)[:len(*c)-1]
	}
	return out
}

func (c *Collection) Size() int {
	return len(*c)
}

func (c *Collection) IsEmpty() bool {
	return false
}
