package main

type book struct {
	id   string
	name string
}

func maintest002() {
	b := book{"eee", "dd"}
	println(true)
	println(b.id)
	s := make([]string, 9)
	s = append(s, "dd")
	println(s[0])
	for i, x := range s {
		println(i)
		println(x)
	}
}
