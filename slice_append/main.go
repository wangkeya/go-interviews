package main

import "fmt"

func main() {
	urls := make(map[string]string, 3)
	urls["baidu"] = "www.baidu.com"
	urls["google"] = "www.google.com"
	urls["csdn"] = "www.csdn.net"

	names := make([]string, len(urls))
	for key, _ := range urls {
		names = append(names, key)
	}
	fmt.Println(names, len(names)) // [   baidu google csdn] 6

	names = make([]string, 0)
	for key, _ := range urls {
		names = append(names, key)
	}
	fmt.Println(names, len(names)) // [google csdn baidu] 3
}
