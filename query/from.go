package query

type FromExpr struct {
	Args []Node
}

type AsExpr struct {
	Args []Node
}

func (f FromExpr) Evaluate(level int) string {
	indent := calcIndent(level)

	result := "\n" + indent + "FROM"
	for i, arg := range f.Args {
		switch arg.(type) {
		case GroupExpr:
			result += arg.Evaluate(level + 1)

		default:
			result += "\n"
			indent = calcIndent(level + 1)
			result += indent
			result += arg.Evaluate(level + 1)
			if i != len(f.Args)-1 {
				result += ",\n"
			}
		}
	}

	return result
}

func (a AsExpr) Evaluate(level int) string {
	result := " AS "

	for _, arg := range a.Args {
		result += arg.Evaluate(level + 1)
	}

	return result
}
