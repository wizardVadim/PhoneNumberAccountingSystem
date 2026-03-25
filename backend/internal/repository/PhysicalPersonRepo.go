package repository


import (
	"database/sql"
	"fmt"
	"phone-accounting-system/internal/models"
)


type PhysicalPersonRepo struct {
	DB *sql.DB
}


func (repo *PhysicalPersonRepo) GetAllPhysicalPersons() []models.PhysicalPerson {
	rows, err := repo.DB.Query("SELECT id, city, person_address, first_name, last_name, second_name, born_year FROM physical_person")
	if err != nil {
		fmt.Println("PhysicalPersonRepo@GetAllPhysicalPersons: Error %w", err)
		return nil
	}
	defer rows.Close()

	list := make([]models.PhysicalPerson, 0, 100)

	for rows.Next() {
		var id int64
		var city string
		var address sql.NullString
		var firstName string
		var lastName string
		var secondName sql.NullString
		var bornYear sql.NullInt16

		if err := rows.Scan(&id, &city, &address, &firstName, &lastName, &secondName, &bornYear); err != nil {
			fmt.Printf("PhysicalPersonRepo@GetAllPhysicalPersons: Error %v\n", err)
			continue
		}

		obj := models.PhysicalPerson{
			Id:         id,
			City:       city,
			Address:    &address.String,
			FirstName:  firstName,
			LastName:   lastName,
			SecondName: &secondName.String,
			BornYear:   &bornYear.Int16,
		}

		if !address.Valid {
			obj.Address = nil
		}
		if !secondName.Valid {
			obj.SecondName = nil
		}
		if !bornYear.Valid {
			obj.BornYear = nil
		}

		list = append(list, obj)
	}

	if err = rows.Err(); err != nil {
		fmt.Printf("PhysicalPersonRepo@GetAllPhysicalPersons: Error %v\n", err)
	}

	return list
}


func (repo *PhysicalPersonRepo) SetPhysicalPerson(person models.PhysicalPerson) error {
	result, err := repo.DB.Exec(`UPDATE physical_person 
		SET 
		city = $2,
		person_address = $3,
		first_name = $4,
		last_name = $5,
		second_name = $6,
		born_year = $7
		WHERE id = $1`,
		person.Id,
		person.City,
		person.Address,
		person.FirstName,
		person.LastName,
		person.SecondName,
		person.BornYear,
	)

	if err != nil {
		fmt.Println("PhysicalPersonRepo@SetPhysicalPerson: Error %w", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("PhysicalPersonRepo@SetPhysicalPerson: Error %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("PhysicalPersonRepo@SetPhysicalPerson: Error didn't find person id %v", person.Id)
	}

	return nil
}


func (repo *PhysicalPersonRepo) CreatePhysicalPerson(person models.PhysicalPerson) (int64, error) {
	var id int64

	err := repo.DB.QueryRow(`INSERT INTO physical_person (city, person_address, first_name, last_name, second_name, born_year) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		person.City,
		person.Address,
		person.FirstName,
		person.LastName,
		person.SecondName,
		person.BornYear,
	).Scan(&id)

	if err != nil {
		fmt.Println("PhysicalPersonRepo@CreatePhysicalPerson: Error %w", err)
		return -1, err
	}

	return id, nil
}


func (repo *PhysicalPersonRepo) RemovePhysicalPerson(person models.PhysicalPerson) error {
	result, err := repo.DB.Exec("DELETE FROM physical_person WHERE id = $1", person.Id)

	if err != nil {
		return fmt.Errorf("PhysicalPersonRepo@RemovePhysicalPerson: Error %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("PhysicalPersonRepo@RemovePhysicalPerson: Error %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("PhysicalPersonRepo@RemovePhysicalPerson: Error didn't find person id %v", person.Id)
	}

	return nil
}


func (repo *PhysicalPersonRepo) GetAllPhysicalPersonsSortedName() []models.PhysicalPerson {
	rows, err := repo.DB.Query(`SELECT id, city, person_address, first_name, last_name, second_name, born_year FROM physical_person
		ORDER BY last_name, first_name, second_name
	`)
	if err != nil {
		fmt.Println("PhysicalPersonRepo@GetAllPhysicalPersonsSortedName: Error %w", err)
		return nil
	}
	defer rows.Close()

	list := make([]models.PhysicalPerson, 0, 100)

	for rows.Next() {
		var id int64
		var city string
		var address sql.NullString
		var firstName string
		var lastName string
		var secondName sql.NullString
		var bornYear sql.NullInt16

		if err := rows.Scan(&id, &city, &address, &firstName, &lastName, &secondName, &bornYear); err != nil {
			fmt.Printf("PhysicalPersonRepo@GetAllPhysicalPersonsSortedName: Error %v\n", err)
			continue
		}

		obj := models.PhysicalPerson{
			Id:         id,
			City:       city,
			Address:    &address.String,
			FirstName:  firstName,
			LastName:   lastName,
			SecondName: &secondName.String,
			BornYear:   &bornYear.Int16,
		}

		if !address.Valid {
			obj.Address = nil
		}
		if !secondName.Valid {
			obj.SecondName = nil
		}
		if !bornYear.Valid {
			obj.BornYear = nil
		}

		list = append(list, obj)
	}

	if err = rows.Err(); err != nil {
		fmt.Printf("PhysicalPersonRepo@GetAllPhysicalPersonsSortedName: Error %v\n", err)
	}

	return list
}


func (repo *PhysicalPersonRepo) GetPhysicalPersonsPhoneNumbersQuantity() map[models.PhysicalPerson]int {

	rows, err := repo.DB.Query(`SELECT 
		COUNT(pn.id) as phone_number_quantity,
		pp.*

		FROM physical_person AS pp
		INNER JOIN phone_number AS pn ON pn.person_id = pp.id

		GROUP BY
		pp.id,
		pp.city,
		pp.person_address,
		pp.first_name,
		pp.last_name,
		pp.second_name,
		pp.born_year
	`)

	if err != nil {
		fmt.Println("PhysicalPersonRepo@GetPhysicalPersonsPhoneNumbersQuantity: Error %w", err)
		return nil
	}
	defer rows.Close()

	list := make(map[models.PhysicalPerson]int)

	for rows.Next() {
		var id int64
		var city string
		var address sql.NullString
		var firstName string
		var lastName string
		var secondName sql.NullString
		var bornYear sql.NullInt16
		var phoneNumberQuantity int

		if err := rows.Scan(&phoneNumberQuantity, &id, &city, &address, &firstName, &lastName, &secondName, &bornYear); err != nil {
			fmt.Printf("PhysicalPersonRepo@GetPhysicalPersonsPhoneNumbersQuantity: Error %v\n", err)
			continue
		}

		obj := models.PhysicalPerson{
			Id:         id,
			City:       city,
			Address:    &address.String,
			FirstName:  firstName,
			LastName:   lastName,
			SecondName: &secondName.String,
			BornYear:   &bornYear.Int16,
		}

		if !address.Valid {
			obj.Address = nil
		}
		if !secondName.Valid {
			obj.SecondName = nil
		}
		if !bornYear.Valid {
			obj.BornYear = nil
		}

		list[obj] = phoneNumberQuantity

	}

	if err = rows.Err(); err != nil {
		fmt.Printf("PhysicalPersonRepo@GetPhysicalPersonsPhoneNumbersQuantity: Error %v\n", err)
	}

	return list

}


func (repo *PhysicalPersonRepo) GetPhysicalPersonById(id int64) *models.PhysicalPerson {
	var personId int64
	var city string
	var address sql.NullString
	var firstName string
	var lastName string
	var secondName sql.NullString
	var bornYear sql.NullInt16

	err := repo.DB.QueryRow(`SELECT id, city, person_address, first_name, last_name, second_name, born_year 
		FROM physical_person WHERE id = $1`, id).Scan(
		&personId, &city, &address, &firstName, &lastName, &secondName, &bornYear)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		fmt.Printf("PhysicalPersonRepo@GetPhysicalPersonById: Error %v\n", err)
		return nil
	}

	obj := &models.PhysicalPerson{
		Id:         personId,
		City:       city,
		Address:    &address.String,
		FirstName:  firstName,
		LastName:   lastName,
		SecondName: &secondName.String,
		BornYear:   &bornYear.Int16,
	}

	if !address.Valid {
		obj.Address = nil
	}
	if !secondName.Valid {
		obj.SecondName = nil
	}
	if !bornYear.Valid {
		obj.BornYear = nil
	}

	return obj
}