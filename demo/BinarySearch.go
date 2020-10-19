package main

import "fmt"

func BinarySearch(array []int,target int,l,r int)int{
	if l>r{
		//出界了，找不到
		return -1
	}

	//从中间开始找
	mid :=(l+r)/2
	middleNum :=array[mid]

	if middleNum ==target{
		return mid //找到了
	}else if middleNum>target{
		//中间的数比目标大，从左边找
		return BinarySearch(array,target,1,mid-1)
	}else {
		// 中间的数比目标还小，从左边找
		return BinarySearch(array,target,mid+1,r)
	}
}

func main()  {

	array:=[]int{1,5,9,81,123,189,333}
	target:=500
	result:=BinarySearch(array,target,0,len(array)-1)

	fmt.Println(target,result)


	target=189
	result = BinarySearch(array,target,0,len(array)-1)
	fmt.Println(target,result)
}