package main

type Ring struct {
	next,pre *Ring     // 前驱和后驱节点
	Value interface{}  // 数据
}

// 初始化空的循环链表，前驱和后驱都指向自己，因为是循环的
func (r *Ring)init()*Ring  {
	r.next = r
	r.pre = r
	return r
}

func main()  {
	r:=new(Ring)
	r.init()
}


// 创建N个节点的循环表
func New(n int)*Ring  {
	if n<=0{
		return nil
	}


	r:=new(Ring)
	p:=r
	for i:=1;i<n;i++{
		p.next=&Ring{pre:p}
		p=p.next
	}

	p.next=r
	r.pre=p
	return r
}

//获取下一个节点
func (r *Ring)Next() *Ring {
	if r.next==nil{
		return r.init()
	}
	return r.next
}

//获取上一个节点
func (r *Ring)Pre()*Ring  {
	if r.pre==nil{
		return r.init()
	}
	return r.pre
}

//获取第n个节点
func (r *Ring)Move(n int)*Ring  {
	if r.next==nil{
		return r.init()
	}

	switch  {
	case n<0:
		for ;n<0;n++{
			r=r.pre
		}
	case n>0:
		for ;n>0;n--{
			r=r.next
		}
	}
	return r
}

// 往节点A，链接一个节点，并且返回之前节点A的后驱节点
func (r *Ring)Link(s *Ring)*Ring  {
	n:=r.Next()
	if s!=nil{
		p:=s.Pre()
		r.next=s
		s.pre=r
		n.pre=p
		p.next=n
	}
	return n
}

//删除节点
func (r *Ring)Unlink(n int)*Ring{
	if n<0{
		return nil
	}
	return r.Link(r.Move(n+1))
}



//查看循环链表的长度

func (r *Ring)Len()int  {
	n:=0
	if r!=nil{
		n=1
		for p:=r.Next();p!=r;p=p.next{
			n++
		}
	}
	return n
}
