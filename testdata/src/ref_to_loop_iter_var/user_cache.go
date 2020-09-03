package ref_to_loop_iter_var

type User struct {
	id   string
	name string
}

var userCache = map[string]*User{}

func cacheUserInfo() {
	users := []User{
		{
			id:   "1",
			name: "1",
		},
		{
			id:   "2",
			name: "2",
		},
		{
			id:   "3",
			name: "3",
		},
		{
			id:   "4",
			name: "4",
		},
	}

	for _, user := range users {
		userCache[user.id] = &user // want "using reference to loop iterator variable"
	}
}
