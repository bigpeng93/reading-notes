package test3

import "fmt"

func Logic() {
	fmt.Println(NewGroup([]string{"ABC123TESTSTRING", "ABC456TESTSTRING", "DEF123TESTSTRING", "DEF456TESTSTRING"}).Process(0, 3, getFilter(3, 6, "123")))
	// 其他的处理过程...
}

// func filterStr(str string) bool {
// 	return str[3:6] == "123"
// }

func getFilter(start, end int, str string) FilterFunc {
	return func(s string) bool {
		return s[start:end] == str
	}
}
