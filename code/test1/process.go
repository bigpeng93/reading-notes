package test1

import "fmt"

func Logic() {
	input := []string{"ABC123TESTSTRING", "ABC456TESTSTRING", "DEF123TESTSTRING", "DEF456TESTSTRING"}

	output := Process(0, 3, input)

	fmt.Println(output)
	// 其他的处理过程...
}

// func Process(intput []string) [][]string {
// 	m := make(map[string][]string)

// 	for _, s := range intput {
// 		h := s[:3]
// 		m[h] = append(m[h], s)
// 	}

// 	output := make([][]string, 0, len(m))
// 	for _, v := range m {
// 		output = append(output, v)
// 	}
// 	return output
// }

func Process(start, end int, intput []string) [][]string {
	m := make(map[string][]string)

	for _, s := range intput {
		h := s[start:end]
		m[h] = append(m[h], s)
	}

	output := make([][]string, 0, len(m))
	for _, v := range m {
		output = append(output, v)
	}
	return output
}
