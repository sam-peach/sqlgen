package query

type GroupExpr struct {
	Args []Node
}

func (g GroupExpr) Evaluate(level int) string {
	result := " ("

	for _, arg := range g.Args {
		nextLevel := level + 1
		indent := calcIndent(nextLevel)

		result += "\n"
		result += indent
		result += arg.Evaluate(nextLevel)
	}

	result += "\n)"

	return result
}
