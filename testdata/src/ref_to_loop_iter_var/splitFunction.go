package ref_to_loop_iter_var

import "fmt"

func splitFunction() {
	var out []*int
	for i := 0; i < 3; i++ {
		innerLoop(i, out) // copy the loop variable because of pass by value
	}
	for _, o := range out {
		fmt.Printf("Value:%d, Address: %p\n", *o, o)
	}
}

func innerLoop(i int, out []*int) {
	fmt.Println(i)
	fmt.Println(&i)
	out = append(out, &i)

	{
		out = append(out, &i)

		{
			out = append(out, &i)
		}
	}
}
