package main

import "fmt"

// 非波那切数列
func F(n int) int {

	if n==0 {
		return 1
	}
	if n==1{
		return 1
	}

	return F(n-1)+F(n-2)

}

func main()  {
	fmt.Println(F(5))
}
