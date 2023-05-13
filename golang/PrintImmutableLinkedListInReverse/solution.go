package PrintImmutableLinkedListInReverse

import "fmt"

/*   Below is the interface for ImmutableListNode, which is already defined for you.
 *
 *   type ImmutableListNode struct {
 *
 *   }
 *
 *   func (this *ImmutableListNode) getNext() ImmutableListNode {
 *		// return the next node.
 *   }
 *
 *   func (this *ImmutableListNode) printValue() {
 *		// print the value of this node.
 *   }
 */

type immutableListNode struct {
	next  *immutableListNode
	value int
}

func (i *immutableListNode) getNext() ImmutableListNode {
	if i == nil {
		return nil
	}
	return i.next
}

func (i *immutableListNode) printValue() {
	if i == nil {
		return
	}
	fmt.Println(i.value)
}

type ImmutableListNode interface {
	getNext() ImmutableListNode
	printValue()
}

func printLinkedListInReverse(head ImmutableListNode) {
	if head == nil {
		return
	}
	printLinkedListInReverse(head.getNext())
	head.printValue()
}
