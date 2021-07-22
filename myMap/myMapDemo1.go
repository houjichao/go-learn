package myMap

import "fmt"

/*
Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。

定义 Map
可以使用内建函数 make 也可以使用 map 关键字来定义 Map:
1.声明变量，默认 map 是 nil
var map_variable map[key_data_type]value_data_type
2.使用 make 函数
map_variable := make(map[key_data_type]value_data_type)
如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
*/
func MapDemo1() {
	fmt.Println("----------map demo1 start----------")

	//创建集合
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	/*
		map 插入 key value
	*/
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India"] = "新德里"
	countryCapitalMap["American"] = "华盛顿"

	/*
		循环输出map
	*/
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是：", countryCapitalMap[country])
	}

	/*
		查看元素在集合中是否存在
	*/
	capital, ok := countryCapitalMap["American"]
	fmt.Println(capital)
	fmt.Println(ok)
	if ok {
		fmt.Println("American的首都是", capital)
	} else {
		fmt.Println("American的首都不存在")
	}

	fmt.Println("----------map demo1 end----------")

}
