package test4

import "strconv"

type Group interface {
	Process(start, end int, filters FilterFunc) [][]string
}

type FilterFunc = func(string) bool

func NewGroup(input interface{}) Group {
	switch in := input.(type) {
	case []string:
		return NewStrGroup(in)
	case []int:
		return NewIntGroup(in)
	}
	return nil
}

type StrGroup struct {
	input []string
}

func NewStrGroup(input []string) *StrGroup {
	return &StrGroup{input: input}
}

func (g *StrGroup) Process(start, end int, filters FilterFunc) [][]string {
	m := make(map[string][]string)

	for _, s := range g.input {
		if filters != nil && !filters(s) {
			continue
		}

		h := s[start:end] //忽略字符串长度检查
		m[h] = append(m[h], s)
	}

	result := make([][]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

type IntGroup struct {
	input []int
}

func NewIntGroup(input []int) *IntGroup {
	return &IntGroup{input: input}
}

func (g *IntGroup) Process(start, end int, filters FilterFunc) [][]string {
	m := make(map[string][]string)

	for _, num := range g.input {
		s := strconv.Itoa(num) //其实这里就是加了个 string to int

		if filters != nil && !filters(s) {
			continue
		}

		h := s[start:end] //忽略字符串长度检查
		m[h] = append(m[h], s)
	}

	result := make([][]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}
