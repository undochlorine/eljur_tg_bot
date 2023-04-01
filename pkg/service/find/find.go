package find

import "errors"

func Interval(tasksBlock [][]string, str string) (int, error) {
	targetStart, targetEnd := "", ""
	facedDash := false
	for _, c := range str {
		if c != '-' && !facedDash {
			targetStart += string(c)
		} else if c == '-' {
			facedDash = true
		} else {
			targetEnd += string(c)
		}
	}
	if !facedDash {
		return -1, errors.New("Invalid inputed string")
	}
	for i, block := range tasksBlock {
		if block[0] == targetStart && block[1] == targetEnd {
			return i, nil
		}
	}
	return -1, nil
}
