package repositories

import (
	"database/sql"
	"turcompany/internal/models"
)

type RoleRepository interface {
	GetByID(id int) (*models.Role, error)
	GetByName(name string) (*models.Role, error)
	List() ([]*models.Role, error)
	Create(role *models.Role) error
	Delete(id int) error
}

type roleRepository struct {
	DB *sql.DB
}

func NewRoleRepository(db *sql.DB) RoleRepository {
	return &roleRepository{DB: db}
}

func (r *roleRepository) GetByID(id int) (*models.Role, error) {
	query := `SELECT id, name, description FROM roles WHERE id = $1`
	role := &models.Role{}
	err := r.DB.QueryRow(query, id).Scan(&role.ID, &role.Name, &role.Description)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (r *roleRepository) GetByName(name string) (*models.Role, error) {
	query := `SELECT id, name, description FROM roles WHERE name = $1`
	role := &models.Role{}
	err := r.DB.QueryRow(query, name).Scan(&role.ID, &role.Name, &role.Description)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (r *roleRepository) List() ([]*models.Role, error) {
	query := `SELECT id, name, description FROM roles ORDER BY id`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []*models.Role
	for rows.Next() {
		r := &models.Role{}
		if err := rows.Scan(&r.ID, &r.Name, &r.Description); err != nil {
			return nil, err
		}
		roles = append(roles, r)
	}
	return roles, nil
}

func (r *roleRepository) Create(role *models.Role) error {
	query := `
		INSERT INTO roles (name, description)
		VALUES ($1, $2)
		RETURNING id
	`
	return r.DB.QueryRow(query, role.Name, role.Description).Scan(&role.ID)
}

func (r *roleRepository) Delete(id int) error {
	query := `DELETE FROM roles WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
