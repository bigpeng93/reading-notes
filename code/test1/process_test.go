package test1

import (
	"fmt"
	"testing"
)

func Test_Process(t *testing.T) {
	input := []string{"ABC123TESTSTRING", "ABC456TESTSTRING", "DEF123TESTSTRING", "DEF456TESTSTRING"}
	output := Process(0, 3, input)
	fmt.Println(output)
}
