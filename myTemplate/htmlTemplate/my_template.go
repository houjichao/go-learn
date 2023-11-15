package main

import (
	"html/template"
	"net/http"
)

const tmpl = `
<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
</head>
<body>
    <h1>{{.Header}}</h1>
    <p>{{.Content}}</p>
</body>
</html>
`

type PageData struct {
	Title   string
	Header  string
	Content string
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "My Page",
		Header:  "Welcome to my website!",
		Content: "This is the main content of the page.",
	}

	t, err := template.New("webpage").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*
*
在这个例子中，我们定义了一个包含三个字段的PageData结构：
Title，Header和Content。HTML模板字符串tmpl使用{{.Title}}，{{.Header}}和{{.Content}}占位符引用这些字段。
我们创建了一个PageData实例，并将其传递给模板执行函数。模板库将占位符替换为相应的字段值，并将结果输出到HTTP响应。

要测试此示例，请将代码保存到名为main.go的文件中，
然后运行go run main.go。在Web浏览器中访问http://localhost:8080，可以看到生成的HTML页面。
*/
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
