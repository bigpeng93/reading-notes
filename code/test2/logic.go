package test2

import "fmt"

func Logic() {
	fmt.Println(NewGroup([]string{"ABC123TESTSTRING", "ABC456TESTSTRING", "DEF123TESTSTRING", "DEF456TESTSTRING"}).Process(0, 3))
	// 其他的处理过程...
}
