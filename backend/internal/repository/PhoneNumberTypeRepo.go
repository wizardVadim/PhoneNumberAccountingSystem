package repository


import (
	"database/sql"
	"fmt"
	"phone-accounting-system/internal/models"
)


type PhoneNumberTypeRepo struct {
	DB *sql.DB
}


func (repo *PhoneNumberTypeRepo) GetAllPhoneNumberTypes() []models.PhoneNumberType {

	rows, err := repo.DB.Query("SELECT id, type_name FROM phone_number_type")

	if err != nil {
		fmt.Println("PhoneNumberTypeRepo@GetAllPhoneNumberTypes: Error %w", err)
		return nil
	}
	defer rows.Close()

	phoneNumberTypeList := make([]models.PhoneNumberType, 0, 10)

	for rows.Next() {

		var typeId int64
		var pntype string

		if err := rows.Scan(&typeId, &pntype); err != nil {
			fmt.Printf("PhoneNumberTypeRepo@GetAllPhoneNumberTypes: Error %v\n", err)
			continue
		}

		pnt := models.PhoneNumberType{Id: typeId, TypeName: pntype}

		phoneNumberTypeList = append(phoneNumberTypeList, pnt)

	}

	if err = rows.Err(); err != nil {
		fmt.Printf("PhoneNumberTypeRepo@GetAllPhoneNumberTypes: Error %v\n", err)
	}

	return phoneNumberTypeList

}


func (repo *PhoneNumberTypeRepo) SetPhoneNumberType(pType models.PhoneNumberType) error {

	result, err := repo.DB.Exec("UPDATE phone_number_type SET type_name = $2 WHERE id = $1", pType.Id, pType.TypeName)

	if err != nil {
		fmt.Println("PhoneNumberTypeRepo@SetPhoneNumberType: Error %w", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("PhoneNumberTypeRepo@SetPhoneNumberType: Error %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("PhoneNumberTypeRepo@SetPhoneNumberType: Error didn't find id %v", pType.Id)
	}

	return nil

}


func (repo *PhoneNumberTypeRepo) CreatePhoneNumberType(pType models.PhoneNumberType) (int64, error) {

	var id int64

	err := repo.DB.QueryRow("INSERT INTO phone_number_type (type_name) VALUES ($1) RETURNING id", pType.TypeName).Scan(&id)

	if err != nil {
		fmt.Println("PhoneNumberTypeRepo@CreatePhoneNumberType: Error %w", err)
		return -1, err
	}

	return id, nil

}


func (repo *PhoneNumberTypeRepo) RemovePhoneNumberType(pType models.PhoneNumberType) error {

	result, err := repo.DB.Exec("DELETE FROM phone_number_type WHERE id = $1", pType.Id)

	if err != nil {
		return fmt.Errorf("PhoneNumberTypeRepo@RemovePhoneNumberType: Error %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("PhoneNumberTypeRepo@RemovePhoneNumberType: Error %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("PhoneNumberTypeRepo@RemovePhoneNumberType: Error didn't find id %v", pType.Id)
	}

	return nil

}


func (repo *PhoneNumberTypeRepo) GetPhoneNumberTypeById(id int64) *models.PhoneNumberType {
	var typeId int64
	var typeName string

	err := repo.DB.QueryRow("SELECT id, type_name FROM phone_number_type WHERE id = $1", id).Scan(&typeId, &typeName)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		fmt.Printf("PhoneNumberTypeRepo@GetPhoneNumberTypeById: Error %v\n", err)
		return nil
	}

	return &models.PhoneNumberType{Id: typeId, TypeName: typeName}
}