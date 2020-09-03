package ref_to_loop_iter_var

import "fmt"

func twoIterationVariables() {
	for i, j := 0, 10; i < j; i++ {
		fmt.Println(i)
		fmt.Println(&i) // want "using reference to loop iterator variable"
		fmt.Println(i)
		fmt.Println(&j) // want "using reference to loop iterator variable"
	}
}
