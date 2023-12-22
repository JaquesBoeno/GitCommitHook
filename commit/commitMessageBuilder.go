package commit

import (
	"fmt"
	"strings"
)

type Value struct {
	Id    string
	Value string
}

func CommitMessageBuilder(template string, values []Value) string {
	s := template

	for _, v := range values {
		str := strings.Split(s, fmt.Sprintf("<%s>", v.Id))

		s = fmt.Sprint(str[0], v.Value, str[1])
	}

	return s
}
