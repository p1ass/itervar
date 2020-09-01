package ref_to_loop_iter_var

import "fmt"

func nestedSlice() {
	var out [][]int
	for _, i := range [][]int{{1}, {2}, {3}} {
		out = append(out, i[:])
	}
	fmt.Println("Values:", out)
}

func nestedArray() {
	var out [][]int
	for _, i := range [][1]int{{1}, {2}, {3}} {
		out = append(out, i[:]) // want "using reference to loop iterator variable"
	}
	fmt.Println("Values:", out)
}
