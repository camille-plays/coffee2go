package main

type person struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Credit int    `json:"credit"`
}

var people = []person{
	{ID: "1", Name: "Suzie", Email: "***REMOVED***", Credit: 0},
	{ID: "2", Name: "Camille", Email: "***REMOVED***", Credit: 0},
	{ID: "3", Name: "Mateusz", Email: "***REMOVED***", Credit: 0},
	{ID: "4", Name: "Steve", Email: "***REMOVED***", Credit: 0},
	{ID: "5", Name: "Kenny", Email: "***REMOVED***", Credit: 0},
	{ID: "6", Name: "Brian", Email: "***REMOVED***", Credit: 0},
}
