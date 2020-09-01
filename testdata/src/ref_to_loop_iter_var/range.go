package ref_to_loop_iter_var

import "fmt"

func rangeLoop() {
	in := []int{1, 2, 3}
	var out []*int
	for _, i := range in {
		fmt.Println(i)
		fmt.Println(&i)       // want "using reference to loop iterator variable"
		out = append(out, &i) // want "using reference to loop iterator variable"

		{
			out = append(out, &i) // want "using reference to loop iterator variable"

			{
				out = append(out, &i) // want "using reference to loop iterator variable"
			}
		}
	}
}
