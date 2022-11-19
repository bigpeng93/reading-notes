package test4

import "fmt"

func Logic() {

	intInput := []int{123456789, 123123678, 456123789, 456456789}
	strInput := []string{"ABC123TESTSTRING", "ABC456TESTSTRING", "DEF123TESTSTRING", "DEF456TESTSTRING"}

	fmt.Println(NewGroup(intInput).Process(0, 3, getFilter(3, 6, "123")))
	fmt.Println(NewGroup(strInput).Process(0, 3, getFilter(3, 6, "123")))
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
