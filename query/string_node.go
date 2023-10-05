package query

type StringNode struct {
	Value string
}

func (s StringNode) Evaluate(level int) string {
	return s.Value
}
