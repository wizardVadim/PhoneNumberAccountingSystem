package repository

import (
	"database/sql"
	"fmt"
	"phone-accounting-system/internal/models"
)

type UserRepo struct {
	DB *sql.DB
}

func (repo *UserRepo) GetAllUsers() []models.User {

	rows, err := repo.DB.Query("SELECT u.id, u.login, u.role_id, ur.role_name FROM \"user\" AS u INNER JOIN user_role AS ur ON u.role_id = ur.id")

	if err != nil {
		fmt.Println("UserRepo@GetAllUsers: Error %w", err)
		return nil
	}
	defer rows.Close()

	userList := make([]models.User, 0, 100)

	for rows.Next() {

		var id int64
		var login string
		var roleId int64
		var role string

		if err := rows.Scan(&id, &login, &roleId, &role); err != nil {
			fmt.Printf("UserRepo@GetAllUsers: Error %v\n", err)
			continue
		}

		u := models.User{Id: id, Login: login, RoleId: roleId, RoleName: role}

		userList = append(userList, u)

	}

	if err = rows.Err(); err != nil {
		fmt.Printf("UserRepo@GetAllUsers: Error %v\n", err)
	}

	return userList

}

func (repo *UserRepo) SetUser(user models.User) error {

	result, err := repo.DB.Exec(`UPDATE user 
		SET 
		login = $2,
		role_id = $3,
		password = crypt($4, gen_salt('bf'))
		WHERE id = $1`,
		user.Id,
		user.Login,
		user.RoleId,
		user.Password,
	)

	if err != nil {
		fmt.Println("UserRepo@SetUser: Error %w", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("UserRepo@SetUser: Error %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("UserRepo@SetUser: Error didn't find user id %v", user.Id)
	}

	return nil

}

func (repo *UserRepo) CreateUser(user models.User) (int64, error) {

	var id int64

	err := repo.DB.QueryRow("INSERT INTO user (login, password, role_id ) VALUES ($1, crypt($2, gen_salt('bf')), $3) RETURNING id",
		user.Login,
		user.Password,
		user.RoleId,
	).Scan(&id)

	if err != nil {
		fmt.Println("UserRepo@CreateUser: Error %w", err)
		return -1, err
	}

	return id, nil

}

func (repo *UserRepo) RemoveUser(user models.User) error {

	result, err := repo.DB.Exec("DELETE FROM user WHERE id = $1", user.Id)

	if err != nil {
		return fmt.Errorf("UserRepo@RemoveUser: Error %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("UserRepo@RemoveUser: Error %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("UserRepo@RemoveUser: Error didn't find user id %v", user.Id)
	}

	return nil

}

func (repo *UserRepo) GetUserById(id int64) *models.User {
	var idVal int64
	var login string
	var password string
	var roleId int64

	err := repo.DB.QueryRow(`SELECT id, login, password, role_id FROM "user" WHERE id = $1`, id).Scan(&idVal, &login, &password, &roleId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		fmt.Printf("UserRepo@GetUserById: Error %v\n", err)
		return nil
	}

	return &models.User{Id: idVal, Login: login, Password: password, RoleId: roleId}
}

func (repo *UserRepo) Auth(user *models.User) bool {
	var idVal int64
	var dbPassword string

	err := repo.DB.QueryRow(`SELECT id, password FROM "user" WHERE login = $1`, user.Login).Scan(&idVal, &dbPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		fmt.Printf("UserRepo@Auth: Error %v\n", err)
		return false
	}

	var isValid bool
	err = repo.DB.QueryRow(`SELECT crypt($1, $2) = $2`, user.Password, dbPassword).Scan(&isValid)
	if err != nil || !isValid {
		return false
	}

	user.Id = idVal
	return true
}
