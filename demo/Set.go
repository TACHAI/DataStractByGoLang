package main

import (
	"fmt"
	"sync"
	"unsafe"
)

// 集合结构体
type Set struct {
	m map[int]struct{}	//用字典来实现，因为字段键不能重复
	len  int 			//集合的大小
	sync.Mutex			//锁，实现并发安全
}

// 添加一个元素
func (s *Set)Add(item int){
	s.Lock()
	defer s.Unlock()

	s.m[item]= struct{}{} //实际往字典添加这个键
	s.len=len(s.m)		  //重新计算元素数量
}


// 移除一个元素
func (s *Set)Remove(item int){
	s.Lock()
	defer s.Unlock()

	//集合没元素直接返回
	if s.len==0{
		return
	}

	delete(s.m,item)//实际从字典删除这个键
	s.len=len(s.m)  //重新计算元素数量
}

//查看是否存在元素
func (s *Set)Has(item int)bool{
	s.Lock()
	defer s.Unlock()
	_,ok:=s.m[item]
	return ok
}

//查看集合大小
func (s *Set)Len()int{
	return s.len
}

//清除集合所有元素
func (s *Set)Clear(){
	s.Lock()
	defer s.Unlock()

	s.m=map[int]struct{}{} //字典重新赋值
	s.len=0				   //大小归零
}

//集合是够空
func (s *Set)IsEmpty()bool{
	if s.Len()==0{
		return true
	}
	return false
}

//将集合转化为列表
func (s *Set)List()[]int{
	s.Lock()
	defer s.Unlock()
	list :=make([]int,0,s.len)
	for item :=range s.m{
		list=append(list,item)
	}
	return list
}


// 为什么使用空结构体
func other(){
	a:= struct {}{}
	b:= struct {}{}
	if a==b{
		fmt.Println("right:%p\n",&a)
	}

	fmt.Println(unsafe.Sizeof(a))
}

func NewSet(cap int64) *Set {
	temp:=make(map[int]struct{},cap)
	return &Set{
		m:temp,
	}
}

func main()  {

	//初始化一个容量为5的不可重复集合
	s:= NewSet(5)

	s.Add(1)
	s.Add(1)
	s.Add(2)

	fmt.Println("list of all items",s.List())

	s.Clear()
	if s.IsEmpty(){
		fmt.Println("2 does exist")
	}

	s.Remove(2)
	s.Remove(3)
	fmt.Println("list of all items",s.List())
}

