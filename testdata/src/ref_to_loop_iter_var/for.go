package ref_to_loop_iter_var

import "fmt"

func forLoop() {
	var out []*int
	for i := 0; i < 3; i++ {
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
