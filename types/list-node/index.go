package listnode

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenListNode(arr []int, pos int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	list := make([]*ListNode, len(arr))

	for i := 0; i < len(arr); i++ {
		list[i] = &ListNode{Val: arr[i]}
		if i > 0 {
			list[i-1].Next = list[i]
		}
	}

	if pos != -1 {
		list[len(list)-1].Next = list[pos]
	}

	return list[0]
}
