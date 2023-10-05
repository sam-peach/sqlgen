package query

import (
	"fmt"
	"testing"
)

func TestSelectEvaluate(t *testing.T) {
	col1 := "my_column"
	col2 := "other_column"
	selectNode := SelectExpr{
		Args: []Node{
			StringNode{
				Value: col1,
			},
			StringNode{
				Value: col2,
			},
		},
	}

	expected := "SELECT\n  my_column,\n  other_column"
	result := selectNode.Evaluate(0)

	Expect(t, expected, result)
}

func Expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		msg := fmt.Sprintf("\nExpected:\n%s\nto equal:\n%s", a, b)
		t.Fatalf(msg)
	}
}
