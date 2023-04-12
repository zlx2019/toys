/**
  @author: Zero
  @date: 2023/4/12 14:10:50
  @desc:

**/

package cond

import "fmt"

func ExampleBool() {
	/*bool*/
	fmt.Println(Bool(false)) //False
	fmt.Println(Bool(true))  // True

	/*int*/
	fmt.Println(Bool(0)) // False
	fmt.Println(Bool(1)) // True

	/*string*/
	fmt.Println(Bool(""))  //False
	fmt.Println(Bool(" ")) //True

	/*Slice*/
	fmt.Println(Bool([]int{}))  //False
	fmt.Println(Bool([]int{1})) //True
}
