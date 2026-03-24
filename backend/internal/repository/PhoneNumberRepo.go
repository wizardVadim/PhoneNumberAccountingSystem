package repository

import (
	"database/sql"
	"fmt"
	"phone-accounting-system/internal/models"
)

type PhoneNumberRepo struct {
	DB *sql.DB
}

func (repo *PhoneNumberRepo) GetAllPhoneNumbers() []models.PhoneNumber {

	rows, err := repo.DB.Query(`SELECT 
		pn.id,
		pn.phone_number_value, 
		pn.person_id, 
		pn.phone_number_type_id, 
		pn.comment, 
		pnt.type_name, 
		pp.first_name, 
		pp.last_name, 
		pp.second_name

		FROM phone_number AS pn 
		INNER JOIN phone_number_type AS pnt ON pn.phone_number_type_id = pnt.id
		INNER JOIN physical_person AS pp ON pn.person_id = pp.id
	`)

	if err != nil {
		fmt.Println("PhoneNumberRepo@GetAllPhoneNumbers: Error %w", err)
		return nil
	}
	defer rows.Close()

	list := make([]models.PhoneNumber, 0, 100)

	for rows.Next() {

		var id int64
		var value string
		var personId int64
		var phoneNumberTypeId int64
		var comment sql.NullString
		var typeName string
		var firstName string
		var lastName string
		var secondName sql.NullString

		if err := rows.Scan(&id,
			&value,
			&personId,
			&phoneNumberTypeId,
			&comment,
			&typeName,
			&firstName,
			&lastName,
			&secondName,
		); err != nil {
			fmt.Printf("PhoneNumberRepo@GetAllPhoneNumbers: Error %v\n", err)
			continue
		}

		obj := models.PhoneNumber{
			Id:                id,
			PhoneNumberValue:  value,
			PersonId:          personId,
			PhoneNumberTypeId: phoneNumberTypeId,
			Comment:           &comment.String,
			PhoneNumberType:   typeName,
			PersonFirstName:   firstName,
			PersonLastName:    lastName,
			PersonSecondName:  &secondName.String,
		}

		if !secondName.Valid {
			obj.PersonSecondName = nil
		}
		if !comment.Valid {
			obj.Comment = nil
		}

		list = append(list, obj)

	}

	if err = rows.Err(); err != nil {
		fmt.Printf("PhoneNumberRepo@PhoneNumberRepo: Error %v\n", err)
	}

	return list

}

func (repo *PhoneNumberRepo) SetPhoneNumber(phone models.PhoneNumber) error {
	result, err := repo.DB.Exec(`UPDATE phone_number 
		SET 
		phone_number_value = $2,
		person_id = $3,
		phone_number_type_id = $4,
		comment = $5
		WHERE id = $1`,
		phone.Id,
		phone.PhoneNumberValue,
		phone.PersonId,
		phone.PhoneNumberTypeId,
		phone.Comment,
	)

	if err != nil {
		fmt.Println("PhoneNumberRepo@SetPhoneNumber: Error %w", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("PhoneNumberRepo@SetPhoneNumber: Error %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("PhoneNumberRepo@SetPhoneNumber: Error didn't find phone number id %v", phone.Id)
	}

	return nil
}

func (repo *PhoneNumberRepo) CreatePhoneNumber(phone models.PhoneNumber) (int64, error) {
	var id int64

	err := repo.DB.QueryRow(`INSERT INTO phone_number (phone_number_value, person_id, phone_number_type_id, comment) 
		VALUES ($1, $2, $3, $4) RETURNING id`,
		phone.PhoneNumberValue,
		phone.PersonId,
		phone.PhoneNumberTypeId,
		phone.Comment,
	).Scan(&id)

	if err != nil {
		fmt.Println("PhoneNumberRepo@CreatePhoneNumber: Error %w", err)
		return -1, err
	}

	return id, nil
}

func (repo *PhoneNumberRepo) RemovePhoneNumber(phone models.PhoneNumber) error {
	result, err := repo.DB.Exec("DELETE FROM phone_number WHERE id = $1", phone.Id)

	if err != nil {
		return fmt.Errorf("PhoneNumberRepo@RemovePhoneNumber: Error %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("PhoneNumberRepo@RemovePhoneNumber: Error %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("PhoneNumberRepo@RemovePhoneNumber: Error didn't find phone number id %v", phone.Id)
	}

	return nil
}

func (repo *PhoneNumberRepo) GetUsersPhoneNumbers(user models.User) []models.PhoneNumber {

	rows, err := repo.DB.Query(`SELECT 
		pn.id,
		pn.phone_number_value, 
		pn.person_id, 
		pn.phone_number_type_id, 
		pn.comment, 
		pnt.type_name, 
		pp.first_name, 
		pp.last_name, 
		pp.second_name

		FROM phone_number AS pn 
		INNER JOIN phone_number_type AS pnt ON pn.phone_number_type_id = pnt.id
		INNER JOIN physical_person AS pp ON pn.person_id = pp.id

		WHERE pp.id = $1`,

		user.Id,

	)

	if err != nil {
		fmt.Println("PhoneNumberRepo@GetAllPhoneNumbers: Error %w", err)
		return nil
	}
	defer rows.Close()

	list := make([]models.PhoneNumber, 0, 100)

	for rows.Next() {

		var id int64
		var value string
		var personId int64
		var phoneNumberTypeId int64
		var comment sql.NullString
		var typeName string
		var firstName string
		var lastName string
		var secondName sql.NullString

		if err := rows.Scan(&id,
			&value,
			&personId,
			&phoneNumberTypeId,
			&comment,
			&typeName,
			&firstName,
			&lastName,
			&secondName,
		); err != nil {
			fmt.Printf("PhoneNumberRepo@GetAllPhoneNumbers: Error %v\n", err)
			continue
		}

		obj := models.PhoneNumber{
			Id:                id,
			PhoneNumberValue:  value,
			PersonId:          personId,
			PhoneNumberTypeId: phoneNumberTypeId,
			Comment:           &comment.String,
			PhoneNumberType:   typeName,
			PersonFirstName:   firstName,
			PersonLastName:    lastName,
			PersonSecondName:  &secondName.String,
		}

		if !secondName.Valid {
			obj.PersonSecondName = nil
		}
		if !comment.Valid {
			obj.Comment = nil
		}

		list = append(list, obj)

	}

	if err = rows.Err(); err != nil {
		fmt.Printf("PhoneNumberRepo@PhoneNumberRepo: Error %v\n", err)
	}

	return list

}
