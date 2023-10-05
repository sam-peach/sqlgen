package query

type SelectExpr struct {
	Args []Node
}

func (s SelectExpr) Evaluate(level int) string {
	result := "SELECT\n"

	for i, arg := range s.Args {
		indent := calcIndent(level + 1)

		result += indent
		result += arg.Evaluate(level + 1)
		if i != len(s.Args)-1 {
			result += ",\n"
		}
	}

	return result
}
