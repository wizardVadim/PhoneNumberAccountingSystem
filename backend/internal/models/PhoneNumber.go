package models

type PhoneNumber struct {
	Id                int64
	PhoneNumberValue  string
	PersonId          int64
	PhoneNumberTypeId int64
	Comment           string
}
