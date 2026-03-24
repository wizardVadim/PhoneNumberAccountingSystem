package models

type PhysicalPerson struct {
	Id         int64	`json:"id"`
	City       string	`json:"city"`
	Address    *string	`json:"address"`
	FirstName  string	`json:"first_name"`
	LastName   string	`json:"last_name"`
	SecondName *string	`json:"second_name"`
	BornYear   *int16	`json:"born_year"`
}
