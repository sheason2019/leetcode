package solution141

// 快慢指针
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fp := head.Next
	sp := head

	for fp != nil {
		if fp == sp {
			return true
		}
		fp = fp.Next
		if fp == nil {
			return false
		}
		fp = fp.Next
		if fp == nil {
			return false
		}
		sp = sp.Next
	}

	return false
}

// Map解法
// func hasCycle(head *ListNode) bool {
// 	attachMap := make(map[*ListNode]bool)
// 	ptr := head
// 	for ptr != nil {
// 		if attachMap[ptr] {
// 			return true
// 		}

// 		attachMap[ptr] = true
// 		ptr = ptr.Next
// 	}

// 	return false
// }
