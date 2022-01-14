package testing

func plus(i int, j int) int64{
	result := i + j
	return int64(result)
}

func TestMe() {
	plus(1, 1)
}