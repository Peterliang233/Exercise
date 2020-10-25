package main

import "fmt"

func main() {
	var s string
	fmt.Scanln(&s)
	l := len(s)
	arr := make([]int, 100010)
	for i := 0; i < l; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	ans := []byte(s)
	for i := 0; i < l; i++ {
		index := arr[i]
		ans[index] = s[i]
	}
	res := string(ans)
	fmt.Println(res)
}