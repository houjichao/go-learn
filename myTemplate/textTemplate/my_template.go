package main

import (
	"bytes"
	"fmt"
	"text/template"
)

const tmpl = `Hello, {{.Name}}! You are {{.Age}} years old.`

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{"John", 30}

	t, err := template.New(tmpl).Parse(tmpl)
	if err != nil {
		panic(err)
	}
	// 当有不存在的模版变量时抛出错误
	t = t.Option("missingkey=error")
	buf := &bytes.Buffer{}
	err = t.Execute(buf, person)
	if err != nil {
		panic(err)
	}
	fmt.Println("-----------", buf.String())
}
