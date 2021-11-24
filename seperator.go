package stringSeperator

import (
	"errors"
)

// Seperator gets a main string and a single char string as surrounder pattern
// and returns all litterals being surrounded by pair of that given char
func Seperator(str string, surrounderChar string) ([]string, error) {

	var counter int
	for _, v := range str {
		if string(v) == surrounderChar {
			counter++
		}
	}

	if counter%2 != 0 {
		return nil, errors.New("serounder templates count most be even")
	}

	literals := make([]string, 0)

	for i := 0; i < len(str); i++ {
		if string(str[i]) == surrounderChar {
			str = str[i+1:]
			lit := ""
			for j := 0; j <= len(str); j++ {
				if string(str[j]) != surrounderChar {
					lit = lit + string(str[j])
				} else {

					literals = append(literals, lit)

					break
				}
				i = j + 1
			}

		}
	}

	return literals, nil

}
