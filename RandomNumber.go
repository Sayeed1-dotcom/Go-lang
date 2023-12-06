package main

import ("fmt" 
       "math/rand")
func main()	{
	fmt.Println(rand.Intn(10000))
	fmt.Println(rand.Intn(1000))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	fmt.Println((rand.Float64() * 12) + 8)
    fmt.Println((rand.Float64() * 9) + 2)
	fmt.Println((rand.Intn(100) + 2) * 3)
	fmt.Println()
}