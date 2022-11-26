package dao

import "github.com/google/uuid"

var Users = []*User{
	{ID: uuid.New().String(), Name: "Suzie", Email: "***REMOVED***", Credit: 0},
	{ID: uuid.New().String(), Name: "Camille", Email: "***REMOVED***", Credit: 0},
	{ID: uuid.New().String(), Name: "Mateusz", Email: "***REMOVED***", Credit: 0},
	{ID: uuid.New().String(), Name: "Steve", Email: "***REMOVED***", Credit: 0},
	{ID: uuid.New().String(), Name: "Kenny", Email: "***REMOVED***", Credit: 0},
	{ID: uuid.New().String(), Name: "Brian", Email: "***REMOVED***", Credit: 0},
}
