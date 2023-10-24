package main

type J interface {
	m(a interface {
		J
	})
}

func main() {
	var j J
	j.m(j)
}
