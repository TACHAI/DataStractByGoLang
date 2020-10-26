package tree

import "fmt"

func BubbleSort(list []int){
	n:=len(list)
	// 在一轮中有没有交换过
	didSwap :=false

	//进行N-1轮迭代
	for i:=n-1;i>0;i--{

		for j:=0;j<i;j++{
			 if list[j]>list[j+1]{
			 	list[j],list[j+1]=list[j+1],list[j]
			 	didSwap=true
			 	}
		}

		// 如果在一轮没有交换过，那么已经排好序了，直接返回
		if !didSwap{
			return
		}
	}
}

func main()  {

	list:=[]int{5,9,1,6,8,14,6,49,25,6,3}
	BubbleSort(list)
	fmt.Println(list)
}
