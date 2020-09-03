package ref_to_loop_iter_var

import "fmt"

func forLoopMap() {
	in := map[int]int{1: 1, 2: 2, 3: 3, 4: 4}
	var out []*int
	for _, value := range in {
		fmt.Println(value)
		fmt.Println(&value)       // want "using reference to loop iterator variable"
		out = append(out, &value) // want "using reference to loop iterator variable"

		{
			out = append(out, &value) // want "using reference to loop iterator variable"

			{
				out = append(out, &value) // want "using reference to loop iterator variable"
			}
		}
	}
}
