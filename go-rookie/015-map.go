package main

import (
	"fmt"
)

func main() {
	//1, 声明map
	var countryCapitalMap map[string]string
	//2, 创建map
	countryCapitalMap = make(map[string]string)
	//3, 初始化map
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	//4, 遍历map
	fmt.Println("遍历countryCapitalMap: ")
	for country := range countryCapitalMap {
		fmt.Printf("country: %s, %s\n", country, countryCapitalMap[country])
	}

	//5, 访问map中的元素
	//country := "United States"
	country := "Japan"
	fmt.Println("")
	fmt.Println("访问countryCapitalMap中的元素: ")
	capital, ok := countryCapitalMap[country]
	if ok {
		fmt.Printf("%s's capital is %s \n", country, capital)
	} else {
		fmt.Printf("%s is not in countryCaptialMap\n", country)
	}

	//6, 删除集合中的元素
	delete(countryCapitalMap, "India")
	fmt.Println("")
	fmt.Println("删除India之后的countryCapitalMap: ")
	for country := range countryCapitalMap {
		fmt.Printf("country: %s, %s\n", country, countryCapitalMap[country])
	}

	//7, map的另一种定义方式
	countryCapitalMap2 := map[string] string {
		"France": "Paris",
		"Italy": "Rome",
		"Japan": "Tokyo", //如果换行, 那么, 最后一行的逗号不能少
	}
	fmt.Println("")
	fmt.Println("另外一种定义map的方式: ")
	for country := range countryCapitalMap2{
		fmt.Printf("country: %s, %s\n", country, countryCapitalMap[country])
	}
}

