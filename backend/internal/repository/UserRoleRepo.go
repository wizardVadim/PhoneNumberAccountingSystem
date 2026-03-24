package repository

import (
	"database/sql"
	"fmt"
	"phone-accounting-system/internal/models"
)

type UserRoleRepo struct {
	DB *sql.DB
}

func (repo *UserRoleRepo) GetAllRoles() []models.UserRole {

	rows, err := repo.DB.Query("SELECT id, role_name FROM user_role")

	if err != nil {
		fmt.Println("UserRoleRepo@GetAllUsers: Error %w", err)
		return nil
	}
	defer rows.Close()

	userRoleList := make([]models.UserRole, 0, 10)

	for rows.Next() {

		var roleId int64
		var role string

		if err := rows.Scan(&roleId, &role); err != nil {
			fmt.Printf("UserRoleRepo@GetAllUsers: Error %v\n", err)
			continue
		}

		ur := models.UserRole{Id: roleId, RoleName: role}

		userRoleList = append(userRoleList, ur)

	}

	if err = rows.Err(); err != nil {
		fmt.Printf("UserRoleRepo@GetAllUsers: Error %v\n", err)
	}

	return userRoleList

}

func (repo *UserRoleRepo) SetUserRole(role models.UserRole) error {

	result, err := repo.DB.Exec("UPDATE user_role SET role_name = $2 WHERE id = $1", role.Id, role.RoleName)

	if err != nil {
		fmt.Println("UserRoleRepo@SetUserRole: Error %w", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("UserRoleRepo@SetUserRole: Error %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("UserRoleRepo@SetUserRole: Error didn't find role id %v", role.Id)
	}

	return nil

}

func (repo *UserRoleRepo) CreateUserRole(role models.UserRole) (int64, error) {

	var id int64

	err := repo.DB.QueryRow("INSERT INTO user_role (role_name) VALUES ($1) RETURNING id", role.RoleName).Scan(&id)

	if err != nil {
		fmt.Println("UserRoleRepo@CreateUserRole: Error %w", err)
		return -1, err
	}

	return id, nil

}

func (repo *UserRoleRepo) RemoveUserRole(role models.UserRole) error {

	result, err := repo.DB.Exec("DELETE FROM user_role WHERE id = $1", role.Id)

	if err != nil {
		return fmt.Errorf("UserRoleRepo@RemoveUserRole: Error %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("UserRoleRepo@RemoveUserRole: Error %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("UserRoleRepo@RemoveUserRole: Error didn't find role id %v", role.Id)
	}

	return nil

}
