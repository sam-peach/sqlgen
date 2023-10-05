package query

const SPACES = 2

func calcIndent(level int) string {
	indent := ""
	for i := 0; i < SPACES*level; i++ {
		indent += " "
	}

	return indent
}
