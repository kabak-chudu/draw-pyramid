package drawpyramid

import "fmt"

func DrawPyramid(val string, rows int) []string {
	strings := []string{}
	for i := 1; i <= rows; i++ {
		res := ""
		//пробел до
		for p := 1; p <= rows-i; p++ {
			res += " "
		}
		//звездочки
		for j := 1; j < i; j++ {
			res += fmt.Sprintf("%s%s", val, val)
		}
		res += fmt.Sprint(val)
		//пробел после
		for p := 1; p <= rows-i; p++ {
			res += " "
		}
		strings = append(strings, res)
	}
	return strings
}
