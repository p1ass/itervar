package ref_to_loop_iter_var

import "fmt"

func copyTheLoopVariable() {
	var out []*int
	for i := 0; i < 3; i++ {
		i := i // Address values changes every time because it copies the loop variable into a new variable.
		out = append(out, &i)
	}
	for _, o := range out {
		fmt.Printf("Value:%d, Address: %p\n", *o, o)
	}
}

func copyTheLoopVariableButInBlockScope() {
	var out []*int
	for i := 0; i < 3; i++ {
		{
			i := i
			out = append(out, &i)
		}
		// Should be detected because this is out of block which copies the loop variable
		out = append(out, &i) // want "using reference to loop iterator variable"
	}
	for _, o := range out {
		fmt.Printf("Value:%d, Address: %p\n", *o, o)
	}
}

func rangeCopyTheLoopVariable() {
	in := []int{1, 2, 3}
	var out []*int
	for _, i := range in {
		i := i // Address values changes every time because it copies the loop variable into a new variable.
		out = append(out, &i)
	}
	for _, o := range out {
		fmt.Printf("Value:%d, Address: %p\n", *o, o)
	}
}

func rangeCopyTheLoopVariableButInBlockScope() {
	in := []int{1, 2, 3}
	var out []*int
	for _, i := range in {
		{
			i := i
			out = append(out, &i)
		}
		// Should be detected because this is out of block which copies the loop variable
		out = append(out, &i) // want "using reference to loop iterator variable"
	}
	for _, o := range out {
		fmt.Printf("Value:%d, Address: %p\n", *o, o)
	}
}
