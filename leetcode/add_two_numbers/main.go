/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// Add a dummy top level node, so we can Next into a loop
	l1 = &ListNode{Next: l1}
	l2 = &ListNode{Next: l2}

	// strings to hold the values
	l1Str := ""
	l2Str := ""

	// Loop through each object and add up its values, doing these
	// seperately since we cannot be gauranteed each object has the same
	// number of levels
	for {
		l1 = l1.Next
		l1Str = strconv.Itoa(l1.Val) + l1Str
		if l1.Next == nil {
			break
		}
	}

	for {
		l2 = l2.Next
		l2Str = strconv.Itoa(l2.Val) + l2Str
		if l2.Next == nil {
			break
		}
	}

	// convert the strings to ints and sum them
	n1 := new(big.Int)
	n2 := new(big.Int)
	sum := new(big.Int)
	u1, _ := n1.SetString(l1Str, 10)
	u2, _ := n2.SetString(l2Str, 10)
	sum = u1.Add(u1, u2)

	// get the length of the string
	sumString := sum.String()
	length := len(sumString) - 1
	reversedStringArr := []string{}

	for i := length; i > 0-1; i-- {
		reversedStringArr = append(reversedStringArr, string(sumString[i]))
	}

	fmt.Println(reversedStringArr)

	result := &ListNode{}
	current := result
	for i, s := range reversedStringArr {
		fmt.Println(i, s)
		current.Val, _ = strconv.Atoi(s)

		if i < length {
			current.Next = &ListNode{}
			current = current.Next
		}
	}

	return result
}
