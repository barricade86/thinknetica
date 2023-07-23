package thing

import (
	"io"
	"reflect"
)

type Employee struct {
	Id   uint
	Rank string
	Name string
	Age  uint
}

func FindCreatureMaxAge(creatures ...interface{}) interface{} {
	var max interface{}
	max = creatures[0]
	for i := 0; i < len(creatures); i++ {
		switch creatures[i].(type) {
		case Customer:
			if creatures[i].(Customer).Age > max.(Customer).Age {
				max = creatures[i]
			}

			continue
		case Employee:
			if creatures[i].(Employee).Age > max.(Employee).Age {
				max = creatures[i]
			}

			continue
		}
	}

	return max
}

func PassOnlyStrings(writer io.Writer, args ...any) []string {
	result := []string{}
	for _, val := range args {
		switch val.(type) {
		case string:
			_, _ = writer.Write([]byte(val.(string)))
			result = append(result, reflect.ValueOf(val).String())
			continue
		default:
			continue
		}
	}

	return result
}
