package main

import "fmt"

func main(){

var Name,Country string 
var Age int
fmt.Print("Enter your Name:")
fmt.Scan(&Name)
fmt.Print("Enter your country name:")
fmt.Scan(&Country)
fmt.Print("Enter your age:")
fmt.Scan(&Age)


fmt.Print("My Name is ", Name)
fmt.Print("\nI am from ", Country)
fmt.Print("\nI am ", Age)

}