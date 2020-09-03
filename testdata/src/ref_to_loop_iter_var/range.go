package ref_to_loop_iter_var

import "fmt"

func rangeLoop() {
	in := []int{1, 2, 3}
	var keyOut []*int
	var valueOut []*int
	for key, value := range in {
		fmt.Println(key)
		fmt.Println(&key)             // want "using reference to loop iterator variable"
		keyOut = append(keyOut, &key) // want "using reference to loop iterator variable"
		fmt.Println(value)
		fmt.Println(&value)                 // want "using reference to loop iterator variable"
		valueOut = append(valueOut, &value) // want "using reference to loop iterator variable"

		{
			valueOut = append(valueOut, &value) // want "using reference to loop iterator variable"

			{
				valueOut = append(valueOut, &value) // want "using reference to loop iterator variable"
			}
		}
	}
	for _, key := range keyOut {
		fmt.Println(*key)
	}
}
