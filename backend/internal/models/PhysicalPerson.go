package models

type PhysicalPerson struct {
	Id         int64
	City       string
	Address    *string
	FirstName  string
	LastName   string
	SecondName *string
	BornYear   *int16
}
