package main

import (
	"github.com/leehom/go-oop/employee"
)

func main()  {
	//fmt.Println(1)
	e := employee.New("leehom","lee",30,20)
	//e := employee.Employee{
	//	FirstName: "leehom",
	//	LastName: "lee",
	//	TotalLeaves: 30,
	//	LeavesTaken: 20,
	//}
	e.LeavesRemaining()
}
