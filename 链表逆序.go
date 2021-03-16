package main


type Node struct {
	Next *Node
	Val  int
}



func main(){

}

// 使得每K个节点之间为一组进行逆序，并且从链表的尾部开始组起，头部剩余节点数量不够一组的不需要逆序。
// 链表:1->2->3->4->5->6->7->8->null, K = 3。那么 6->7->8，3->4->5，1->2各位一组。调整后：1->2->5->4->3->8->7->6->null。其中 1，2不调整，因为不够一组。



// 反转链表
func reverseList(head *Node)(*Node){

	if head == nil || head.Next == nil  {
		return head
	}

	newList := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newList
}


//你只需要先把单链表进行一次逆序，逆序之后就能转化为从头部开始组起了，
// 然后按照我上面的解法，处理完之后，把结果再次逆序即搞定。两次逆序相当于没逆序。
func slove(head *Node,k int )*Node{

	head = reverseList(head)

	head = reverseGroup(head,k)

	head = reverseList(head)

	return head
}


func reverseGroup(head *Node,k int)(*Node){

	temp := head

	//找个分组节点
	for i:=1; i<k && temp != nil ;i++{
		temp = temp.Next
	}
	//判断节点的数量是否能够凑成一组
	if temp == nil {
		return head
	}
	h := temp.Next
	temp.Next = nil

	//把当前的组进行逆序
	newHead := reverseList(head)
	//把之后的节点进行分组逆序
	newTemp := reverseGroup(h,k)
	// 把两部分连接起来
	head.Next = newTemp
	return newHead
}



func reverseList1(head *Node)(*Node){

	if head == nil || head.Next == nil  {
		return head
	}

	head.Next.Next = head
}