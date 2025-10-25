package main

import "fmt"

type Stack struct {
	Data []int
}

func (s *Stack) PopBack() int {
	if len(s.Data) == 0 {
		fmt.Println("Error!!! Stack is empty. Returns 0")
		return 0
	} else {
		res := s.Data[len(s.Data)-1]
		s.Data = s.Data[:len(s.Data)-1]
		return res
	}
}

func (s *Stack) PopFront() int {
	if len(s.Data) == 0 {
		fmt.Println("Error!!! Stack is empty. Returns 0")
		return 0
	} else {
		res := s.Data[0]
		s.Data = s.Data[1:]
		return res
	}
}

func (s *Stack) PushBack(x int) {
	s.Data = append(s.Data, x)
}

func (s *Stack) PushFront(x int) {
	s.Data = append(s.Data, 0)
	for i := len(s.Data) - 1; i > 0; i-- {
		s.Data[i] = s.Data[i-1]
	}
	s.Data[0] = x
}

func (s *Stack) IsEmpty() bool {
	return len(s.Data) == 0
}

func (s *Stack) Size() int {
	return len(s.Data)
}

func (s *Stack) Clear() {
	s.Data = nil
}

func (s *Stack) fillStack(n int) {
	fmt.Printf("Input %d numbers: ", n)
	for i := 0; i < n; i++ {
		var temp int
		fmt.Scan(&temp)
		s.Data = append(s.Data, temp)
	}
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BST struct {
	Root *Node
}

func (n *BST) Insert(x int) {
	if n.Root == nil {
		n.Root = &Node{x, nil, nil}
	} else {
		n.Root = insert(n.Root, x)
	}
}

func insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{value, nil, nil}
	} else if value < node.Value {
		node.Left = insert(node.Left, value)
	} else if value > node.Value {
		node.Right = insert(node.Right, value)
	}
	return node
}

func (n *BST) Find(value int) bool {
	if n.Root == nil {
		return false
	}
	return find(n.Root, value)
}

func find(node *Node, value int) bool {
	if node == nil {
		return false
	}
	if node.Value == value {
		return true
	} else if value < node.Value {
		return find(node.Left, value)
	} else {
		return find(node.Right, value)
	}

}

func (n *BST) Depth() int {
	if n.Root == nil {
		return 0
	}
	return depth(n.Root)
}

func depth(node *Node) int {
	if node == nil {
		return 0
	}
	left := depth(node.Left)
	right := depth(node.Right)
	return 1 + max(left, right)
}

func (n *BST) Remove(value int) {
	if n.Root == nil {
		fmt.Println("Tree is empty")
		return
	}
	n.Root = remove(n.Root, value)
}

func remove(node *Node, value int) *Node {
	if node == nil {
		fmt.Println("Elemnt not found")
		return nil
	}
	if node.Value == value {
		if node.Left == nil && node.Right == nil {
			return nil
		} else if node.Left == nil && node.Right != nil {
			return node.Right
		} else if node.Left != nil && node.Right == nil {
			return node.Left
		} else {
			successor := findSuccessor(node.Right)
			node.Value = successor.Value
			node.Right = remove(node.Right, successor.Value)
			return node
		}
	} else {
		if value < node.Value {
			node.Left = remove(node.Left, value)
		} else {
			node.Right = remove(node.Right, value)
		}
		return node
	}
}

func findSuccessor(node *Node) *Node {
	if node.Left != nil {
		return findSuccessor(node.Left)
	}
	return node
}

func main() {
	var check string
	fmt.Printf("If you want to test BST, enter b, else, enter s: ")
	fmt.Scan(&check)
	if check == "s" {
		s := Stack{}
		var length int
		var add_number int
		var get_number int
		var sizeOF int

		fmt.Printf("Input a length of slice: ")
		fmt.Scan(&length)
		if length != 0 {
			s.fillStack(length)
		}

		fmt.Println()
		fmt.Println("Your stack: ", s)
		fmt.Println()

		fmt.Printf("Func PushFront(). Input a number to add to your stack: ")
		fmt.Scan(&add_number)
		s.PushFront(add_number)
		fmt.Println("Your stack: ", s)
		fmt.Println()

		fmt.Printf("Func PushBack(). Input a number to add to your stack: ")
		fmt.Scan(&add_number)
		s.PushBack(add_number)

		fmt.Println("Your stack after func PushBack() : ", s)
		fmt.Println()

		fmt.Println("Func PopBack()")
		get_number = s.PopBack()
		fmt.Printf("The last value of stack is - %d\n", get_number)

		fmt.Println("Your stack after func PopBack() : ", s)
		fmt.Println()

		fmt.Println("Func PopFront()")
		get_number = s.PopFront()
		fmt.Printf("The first value of stack is - %d\n", get_number)

		fmt.Println("Your stack after func PopFront() : ", s)
		fmt.Println()

		b := s.IsEmpty()
		fmt.Printf("Func IsEmpty() - %t\n", b)
		fmt.Println()

		sizeOF = s.Size()
		fmt.Println("Func Size(). Size of your stack is ", sizeOF)
		fmt.Println()

		fmt.Println("Func Clear()")
		fmt.Println("Your stack before Clear() ", s)
		s.Clear()
		fmt.Println("Your stack after Clear() ", s)
	} else {
		b := BST{}
		fmt.Printf("Input number of numbers to insert: ")
		var n int
		fmt.Scan(&n)
		for i := 0; i < n; i++ {
			fmt.Printf("Input a number (%d/%d) : ", i+1, n)
			var temp int
			fmt.Scan(&temp)
			b.Insert(temp)
		}

		fmt.Printf("Input a number to check it in tree : ")
		fmt.Scan(&n)
		fmt.Printf("Is %d in tree? %t\n", n, b.Find(n))
		fmt.Printf("Input another number to check it in tree : ")
		fmt.Scan(&n)
		fmt.Printf("Is %d in tree? %t\n", n, b.Find(n))

		fmt.Println()
		fmt.Println("Depth of tree is ", b.Depth())

		fmt.Printf("Input a number to remove it from tree : ")
		fmt.Scan(&n)
		b.Remove(n)

		fmt.Println("Lets check the depth again - ", b.Depth())
	}

}
