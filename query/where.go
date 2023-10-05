package query

type WhereExpr struct {
	Args []Node
}

type AndExpr struct {
	Args []Node
}

type OrExpr struct {
	Args []Node
}

func (w WhereExpr) Evaluate(level int) string {
	indent := calcIndent(level)

	result := "\n" + indent + "WHERE\n"

	for i, arg := range w.Args {
		indent := calcIndent(level + 1)

		result += indent
		result += arg.Evaluate(level + 1)
		if i != len(w.Args)-1 {
			result += ",\n"
		}
	}

	return result
}

func (a AndExpr) Evaluate(level int) string {
	indent := calcIndent(level + 1)
	result := "\n" + indent + "AND"
	for i, arg := range a.Args {

		result += " "
		result += arg.Evaluate(level + 1)
		if i != len(a.Args)-1 {
			result += ",\n"
		}
	}

	return result
}

func (a OrExpr) Evaluate(level int) string {
	indent := calcIndent(level + 1)
	result := "\n" + indent + "OR"
	for i, arg := range a.Args {

		result += " "
		result += arg.Evaluate(level + 1)
		if i != len(a.Args)-1 {
			result += ",\n"
		}
	}

	return result
}
