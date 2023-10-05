package query

type Node interface {
	Evaluate(int) string
}

type Query struct {
	Args []Node
}

func Select(cols ...string) *Query {
	stringArgs := make([]Node, len(cols))
	for i, col := range cols {
		stringArgs[i] = StringNode{Value: col}
	}

	mainArgs := []Node{SelectExpr{
		Args: stringArgs,
	}}

	return &Query{
		Args: mainArgs,
	}
}

func (q *Query) From(from interface{}) *Query {
	var arg Node

	switch from.(type) {
	case string:
		if value, ok := from.(string); ok {
			arg = FromExpr{
				Args: []Node{
					StringNode{
						Value: value,
					},
				},
			}
		}
	case *Query:
		if value, ok := from.(*Query); ok {
			arg = FromExpr{
				Args: []Node{
					GroupExpr{
						Args: []Node{
							value,
						},
					},
				},
			}
		}
	}

	q.Args = append(q.Args, arg)

	return q
}

func (q *Query) As(label string) *Query {
	q.Args = append(q.Args, AsExpr{
		Args: []Node{
			StringNode{
				Value: label,
			},
		},
	})

	return q
}

func (q *Query) Where(clause string) *Query {
	q.Args = append(
		q.Args,
		WhereExpr{
			Args: []Node{
				StringNode{
					Value: clause,
				},
			},
		},
	)

	return q
}

func (q *Query) And(clause string) *Query {
	q.Args = append(
		q.Args,
		AndExpr{
			Args: []Node{
				StringNode{
					Value: clause,
				},
			},
		},
	)

	return q
}

func (q *Query) Or(clause string) *Query {
	q.Args = append(
		q.Args,
		OrExpr{
			Args: []Node{
				StringNode{
					Value: clause,
				},
			},
		},
	)

	return q
}

func (q Query) Evaluate(level int) string {
	result := ""

	for _, arg := range q.Args {
		result += arg.Evaluate(level)
	}

	return result
}

func (q Query) ToSql() string {
	result := ""

	for _, arg := range q.Args {
		result += arg.Evaluate(0)
	}

	return result
}
