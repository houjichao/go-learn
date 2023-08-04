package main

import "fmt"

func main() {
	status := 32
	statusP := &status
	where := map[string]any{
		"belongModule": 1,
		"cid":          "2222",
	}
	if statusP != nil {
		where["status"] = *statusP
	}

	fmt.Println(where)
}
