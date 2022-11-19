# 代码抽象的过程

结合工作以来写代码踩过的坑，用一个简单的例子来说明一段代码的进化过程：

-  面向过程编程 -> 面向对象编程 -> 函数式编程 -> 抽象

我就以golang为例说明一下，为什么使用golang，而不是其他语言：
首先 java 语言自带面向对象属性，就是一个function必须依托它所在的class，在面向对象的时候，不好说明；python 我是真的不熟，

## 一、要解决的问题

现在描述一下要解决的问题：

```
给一串字符串，根据一个字符串前三位进行分组。

input: []string{"ABC123TESTSTRING","ABC456TESTSTRING","DEF123TESTSTRING","DEF456TESTSTRING"}

output: []string{"ABC123TESTSTRING","ABC456TESTSTRING"} 和 []string{"DEF123TESTSTRING","DEF456TESTSTRING"}

注:
    1、output 输出的数组不区顺序
    2、假设输入的字符串的长度完全满足处理要求，不需要进行一些长度的校验
```

根据这个问题的输入输出，上层的调用方可能还有其他逻辑，我这里就写的简单一点，简单的打印一下：

```golang
func Logic() {
	input := []string{"ABC123TESTSTRING", "ABC456TESTSTRING", "DEF123TESTSTRING", "DEF456TESTSTRING"}

    // 具体处理的过程
    resoult := Process(input)
	
	// 后续的处理过程...
	fmt.Println(output)
}
```

接下来我们就一步一步的填充这个具体处理过程

## 二、代码优化的过程

### 2.1 面向过程编程

当我们拿到这个问题的时候，我们一般首先想到的是肯定是面向过程编程，代码很简单，如下：

```golang
func Logic() {
	input := []string{"ABC123TESTSTRING", "ABC456TESTSTRING", "DEF123TESTSTRING", "DEF456TESTSTRING"}

    // 具体处理的过程
	m := make(map[string][]string)

	for _, s := range intput {
		h := s[:3]
		m[h] = append(m[h], s)
	}

	output := make([][]string, 0, len(m))
	for _, v := range m {
		output = append(output, v)
	}
	
	// 后续的处理过程...
	fmt.Println(output)
}
```

写完这段代码，我们发现中间的处理过程，其实就只是针对input的，然后得到一个叫result的数组，期间就没有引用其他变量了，所以我们就可以将这段代码抽出来，
单独作为一个function，然后就成了这个：
```golang
func Process(input []string)[][]string{
	m := make(map[string][]string)

	for _, s := range intput {
		h := s[:3]
		m[h] = append(m[h], s)
	}

	result := make([][]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
    return result
}
```
我们现在封装出来一个根据字符串**前三个字符**进行分组的函数，所以现在Process函数封装的并不通用，它最好做的能根据字符串的任意起始位置来进行分组，所以我们现在重新改写一下Process函数。

最终的代码就成了：

```golang
func Logic() {
	input := []string{"ABC123TESTSTRING", "ABC456TESTSTRING", "DEF123TESTSTRING", "DEF456TESTSTRING"}

	output := Process(0, 3, input)

	fmt.Println(output)
	// 其他的处理过程...
}

func Process(start, end int, intput []string) [][]string {
	m := make(map[string][]string)

	for _, s := range intput {
		h := s[start:end]
		m[h] = append(m[h], s)
	}

	result := make([][]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

```

注：
    现在函数被拆成了两个短小精悍的小function，就不会导致将所有的处理逻辑都放在Logic函数中，如果Logic还有其他逻辑，Logic函数就会越来长，越来越复杂。
    这种代码其实都不应该算是什么优化，这应该是在大脑中的一个潜意识，在最开始写的时候，就应该做的。

在这个Process函数中，你会发现，它处理的数据仅仅就是输入的input数组，现在我们转换一下思路，这个Process函数是input这个数组的一个行为，想到这就进入面向对象的编程过程

[代码地址](./test1/)

## 2.2 面向对象编程

我们认为Process分组是input 数组做出的一个动作，这样认为的依据就是process处理的数据是input它自己。

那我们就定义一个新的类型，并实现Process函数，现在代码如下

```golang
type Group struct {
	input []string
}

func NewGroup(input []string) *Group {
	return &Group{input: input}
}

func (g *Group) Process(start, end int) [][]string {
	m := make(map[string][]string)

	for _, s := range g.input {
		h := s[start:end] //忽略字符串长度检查
		m[h] = append(m[h], s)
	}

	result := make([][]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}
```
注：暂且叫Group吧，这个应该根据业务中具体含义来定义，最好是名词而不是动词。

现在上层的Logic函数就可以这样来写了：

```golang
func Logic() {
	input := []string{"ABC123TESTSTRING", "ABC456TESTSTRING", "DEF123TESTSTRING", "DEF456TESTSTRING"}

	output := NewGroup(input).Process(0, 3)

	fmt.Println(output)
	// 其他的处理过程...
}
```
现在Logic函数甚至可以一行代码解决：

```golang
func Logic() {
	fmt.Println(NewGroup([]string{"ABC123TESTSTRING", "ABC456TESTSTRING", "DEF123TESTSTRING", "DEF456TESTSTRING"}).Process(0, 3))
	// 其他的处理过程...
}
```
    注：

    因为go语言中，如果属于结构体的函数，会在函数前面加上结构的引用，那么这个结构体其实也可以放到function 后面作为输入参数，所以可以根据这个来判断是不是该抽象了。
    要仔细分析输入的所有的参数，在输入的参数当中主要是分为两类，一个是实体的类型，java和python中的class，golang中的struct；另一种就是执行具体操作的类型，比如读取数据库的client，http的client，或者内部执行业务逻辑controller。我们抽象出来的类型最好是指包含一种实体类型，不要包含其它的业务含义的类型进来，如果是特意根据业务进行的组合，这个无所谓，其它的参数就像实例中的start和end一样，放在function的请求参数中。而一些执行具体操作的类型，这些在抽象到使用完，交给垃圾回收销毁之后，并不会发生变化，可以在类型实例化时候，就将其赋值进去。

[代码地址](./test2/)

### 2.3 函数式编程

函数式编程，说简单点就是将函数作为另一个函数请求的入参，在另一个函数执行的时候，会根据自己的逻辑执行输入的函数。

我现在临时加个需求：

```
要过滤掉 3-5 位不是"123"的字符串，并进行分组。
```

首先我们在Logic函数中现将input数组过滤一遍，然后再按照刚才的逻辑进行处理，但是这样操作就会将Logic函数的代码的复杂性增加。在软件开发的时候，有个原则就是要“封装复杂性”，就是将代码的复杂性尽量封装到底层的代码中，不要去扩大上层代码的复杂性。

我们首先定义一个function类型 FilterFunc ，用于执行数据过滤

那么我们就可以在 Group的Process的参数中加一个FilterFunc类型参数，在执行分组的时候，将字符进行过滤。

现在Group就变成了这样的

```golang
type Group struct {
	input []string
}

func NewGroup(input []string) *Group {
	return &Group{input: input}
}

func (g *Group) Process(start, end int, filters FilterFunc) [][]string {
	m := make(map[string][]string)

	for _, s := range g.input {
		if filters != nil && filters(s) {
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

func Logic() {
	fmt.Println(NewGroup([]string{"ABC123TESTSTRING", "ABC456TESTSTRING", "DEF123TESTSTRING", "DEF456TESTSTRING"}).Process(0, 3, filterStr))
	// 其他的处理过程...
}

func filterStr(str string) bool {
	return str[3:6] == "123"
}
```

我发现filter函数里面的`3,6,"123"`也是hardcode的，我又可以用函数式编程变得通用一点,filter就变成这样了。

```golang
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
```

代码到这里你会发现，在Group中和getFilter函数中没有一个hardcode的参数，代码中的参数全部都在Logic当中。


	注：
	关于这个FilterFunc的定义，应该和Group这个struct定义在一起，而不能随处定义，Group可以提供默认的Filter，上层业务也可以自定义。
	Filter可以使用golang的可变参数，也可以使用 golang里的option模式。

[代码地址](./test3/)

### 2.4 抽象

我现在又要加需求

```
在之前的基础上，我想对于一系列的int数组也按照上面的逻辑进行处理。为了处理方便，这次返回仍旧是 string 数组的分组
```

我现在安装之前Group的做法，在新定义一个类型，不过这次里面的参数是int数组，而不是字符串数组。为了和int区分开，我们将之前 Group struct 改名字为 StrStruct，
同时新建一个IntStruct，用来专门处理 int类型的数组。

此时，

```golang
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
```
我们发现两个struct中有想用的Process函数，而且出参和入参相同，所以我们就可以抽象出一个 interface，代指具有相同行为的一类，
同时创建一个NewGroup函数，类似与一个工厂模式，根据输入参数的不同，返回不同的interface实例。

我特别讨厌在写代码的时候就考虑设计模式，就成了单纯的套用模板而将代码写得啥都不是，设计模式一种解决方式，在写代码的时候就已经不知不觉地在使用。

```golang
type Group interface {
	Process(start, end int, filters FilterFunc) [][]string
}

func NewGroup(input interface{}) Group {
	switch in := input.(type) {
	case []string:
		return NewStrGroup(in)
	case []int:
		return NewIntGroup(in)
	}
	return nil
}
```

logic 代码

```golang
func Logic() {

	intInput := []int{123456789, 123123678, 456123789, 456456789}
	strInput := []string{"ABC123TESTSTRING", "ABC456TESTSTRING", "DEF123TESTSTRING", "DEF456TESTSTRING"}

	fmt.Println(NewGroup(intInput).Process(0, 3, getFilter(3, 6, "123")))
	fmt.Println(NewGroup(strInput).Process(0, 3, getFilter(3, 6, "123")))
	// 其他的处理过程...
}

func getFilter(start, end int, str string) FilterFunc {
	return func(s string) bool {
		return s[start:end] == str
	}
}
```

抽象是一种精确的表达......

[代码地址](./test4/)