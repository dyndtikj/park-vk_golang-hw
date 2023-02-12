package main

import "fmt"
import "calculator/stack"

func main() {
	fmt.Println("Calculator")
	st := stack.New()
	st.Push(1)
	fmt.Println(st.Top())
	st.Push(2)
	fmt.Println(st.Top())
	st.Push(3)
	fmt.Println(st.Top())
	st.Push(4)
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
	fmt.Println(st.Top())

}
