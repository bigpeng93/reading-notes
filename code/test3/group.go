package test3

type Group struct {
	input []string
}

func NewGroup(input []string) *Group {
	return &Group{input: input}
}

func (g *Group) Process(start, end int, filters FilterFunc) [][]string {
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

type FilterFunc = func(string) bool
