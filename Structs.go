package main

import ("log")

var s = "One"
func main(){
	var s2 = "Two"
   log.Println("s is", s)
   log.Println("s2 is",s2)
   say("Hello","Golang", "Hello Golang")
}
func say(s string,s3 string,s4 string )(string,string,string){
	log.Println("s from say func is", s)
	log.Println("s3 from say func is",s3)
	log.Println("s4 from say func is",s4)


	return s,s3,s4
}