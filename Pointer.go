package main

import "log"

func main(){
	var Game string
	Game ="Cricket"
	log.Println("My favourite game is", Game)

	changeGame(&Game)
	log.Println("After func call my favourite game is", Game)

	var ID int
	ID = 218801014
	log.Println("My student id is", ID)

	changeID(&ID)
	log.Println("After func call my new ID is", ID)


}
    func changeGame(s *string){
		newGame :="Football"
		*s = newGame
	}
	func changeID(s *int){
	    newID := 14
		*s = newID
	}