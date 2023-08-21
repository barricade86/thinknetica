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
	if len(creatures) == 0 {
		return 0
	}

	max = creatures[0]
	for i := 0; i < len(creatures); i++ {
		switch person := creatures[i].(type) {
		case Customer:
			if person.Age > max.(Customer).Age {
				max = creatures[i]
			}

			continue
		case Employee:
			if person.Age > max.(Employee).Age {
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
