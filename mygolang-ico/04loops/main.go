package main

func main() {
	//loops in go
	//for loop
	for i := 0; i < 5; i++ {
		println(i)
	}
	//for loop with range
	nums := []int{2, 3, 4}
	for index, value := range nums {
		println(index, value)
	}
	//while loop
	i := 1
	for i <= 5 {
		println(i)
		i++
	}

}
