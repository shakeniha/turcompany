package repositories
import (
	"database/sql"
	"fmt"
	"turcompany/internal/models"
)
type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}
func (r *UserRepository) CreateUser(user *models.User)error{
	query:= `INSERT INTO users (company_name, bin_iin, email, password_hash, role_id)
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`
	return r.db.QueryRow(query, user.CompanyName, user.BinIin, user.Email, user.PasswordHash, user.RoleID).Scan(&user.ID)
}
func 