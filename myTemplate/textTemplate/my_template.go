package main

import (
	"bytes"
	"fmt"
	"text/template"
)

const tmpl = `Hello, {{.Name}}! You are {{.Age}} years old.Your school is {{.School}}`

// 使用{{if .City}} Your city is {{.City}}.{{end}}检查City字段是否存在。由于Person结构中没有City字段，此部分将被忽略，不会导致错误。程序将输出：
const tmpl1 = `Hello, {{.Name}}! You are {{.Age}} years old.{{if .City}} Your city is {{.City}}.{{end}}`

type Person struct {
	Name string
	Age  int
}

func main() {
	//person := Person{"John", 30}

	t, err := template.New(tmpl).Parse(tmpl)
	if err != nil {
		panic(err)
	}

	personMap := map[string]interface{}{
		"Name": "Alice",
		"Age":  30,
	}

	// 当有不存在的模版变量时抛出错误
	/*t = t.Option("missingkey=zero")
	buf := &bytes.Buffer{}
	err = t.Execute(buf, person)
	if err != nil {
		panic(err)
	}
	fmt.Println("-----------", buf.String())*/

	/**
	Option方法 它接受一个或多个字符串参数，这些字符串参数表示要设置的选项。目前支持以下选项：
	missingkey=zero：当模板执行期间访问映射时，如果映射中不存在请求的键，则将该键的值视为零值。这是默认行为。
	missingkey=error：如果映射中不存在请求的键，则模板执行将返回错误。
	missingkey=invalid：如果映射中不存在请求的键，模板执行将产生无效的操作。
	*/
	/**
	missingkey=zero选项只适用于映射（map），而不是结构体。在结构体中，访问不存在的字段会导致编译错误。
	*/
	//如果您需要在模板中处理不存在的字段，建议使用映射（map）而不是结构体。
	//这样，您可以利用missingkey=zero选项来处理缺失的键。以下是一个使用映射的示例：
	t = t.Option("missingkey=zero")
	buf := &bytes.Buffer{}
	err = t.Execute(buf, personMap)
	if err != nil {
		panic(err)
	}
	fmt.Println("***********", buf.String())
}
