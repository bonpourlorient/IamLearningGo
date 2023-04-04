package main

import (
	"encoding/json"
	"fmt"
	"os"
)



type Email struct {
	ID      int
	Kind    string
	Address string
}

type Interest struct {
	ID   int
	Name string
}

type Person struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Gender    string
	Email     []Email
	Interest  []Interest
}

func CreatePerson(ID int,FirstName string,LastName string,UserName string,Gender string , Email []Email,Interest []Interest) Person {

	return Person {
		ID : ID,
		FirstName: FirstName,
		LastName: LastName,
		UserName: UserName,
		Gender: Gender,
		Email: Email,
		Interest: Interest,
	}
}

func GetPerson(p *Person) string {
	return p.FirstName + " " + p.LastName
}

func GetPersonEmailAddr(p *Person, i int) string {

	return p.Email[i].Address
}

func GetPersonEmail(p *Person, i int) Email {
	return p.Email[i]
}

func WriteMessage(msg string) {
	fmt.Println(msg)
}
func WriteStarLine() {
	fmt.Println("*******************")
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error!")
		os.Exit(1)
	}
}

func saveJSON(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	defer outFile.Close()
	checkError(err)
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)

}

func main() {
	
	
	person := CreatePerson(1,"James","Snow","hy3n4","E",[]Email{ Email{1,"work","kancad@nomail.com"},Email{2,"fun","hy3n4@nomail.com"}},[]Interest {Interest{1,"Go"},Interest{2,"Python"},Interest{3,"Go"}})

	WriteMessage("Reading Operation Started")
	WriteMessage("Personal Fullname")
	WriteStarLine()
	res := GetPerson(&person)
	WriteMessage(res)
	WriteStarLine()
	WriteMessage("\n")

	WriteMessage("Personal email With index")
	WriteStarLine()

	resEmail := GetPersonEmailAddr(&person, 1)
	WriteMessage(resEmail)
	WriteStarLine()

	WriteMessage("\n")

	WriteMessage("Personal Email Object With Index")
	WriteStarLine()
	resEmailObject := GetPersonEmail(&person, 1)
	fmt.Println(resEmailObject)

	WriteStarLine()

	WriteMessage("Writing Operation Started")

	saveJSON("person.json", person)

	WriteMessage("Operation ended")
}
