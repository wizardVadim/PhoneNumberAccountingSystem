package models

type PhoneNumber struct {
	Id                int64  `json:"id"`
	PhoneNumberValue  string `json:"phone_number_value"`
	PersonId          int64  `json:"person_id"`
	PersonFirstName   string `json:"person_first_name"`
	PersonLastName    string `json:"person_last_name"`
	PersonSecondName  *string `json:"person_second_name"`
	PhoneNumberTypeId int64  `json:"phone_number_type_id"`
	PhoneNumberType   string `json:"phone_number_type"`
	Comment           *string `json:"comment"`
}
