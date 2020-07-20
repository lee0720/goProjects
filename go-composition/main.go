package main

import "fmt"

type author struct {
	firstName string
	lastName string
	bio 	string
}

type post struct {
	title 		string
	content 	string
	author
}

func (a author) fullName() string  {
	return fmt.Sprintf("%s %s",a.firstName,a.lastName)
}

func (p post) details()  {
	fmt.Println("Title: ",p.title)
	fmt.Println("Content: ",p.content)
	fmt.Println("Author: ",p.author.fullName())
	fmt.Println("Bio: ",p.author.bio)

}

func main()  {
	aut:= author{
		"leehom","lee","Golang Enthusiast",
	}
	pos := post{"how to learn golang","just do it",aut}
	pos.details()
}


