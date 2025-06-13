package repositories

import (
	"database/sql"
	"turcompany/internal/models"
)

type LeadRepository struct {
	db *sql.DB
}

func NewLeadRepository(db *sql.DB) *LeadRepository {
	return &LeadRepository{db: db}
}
func (r *LeadRepository) Create(lead *models.Leads) error {

	query := `
		INSERT INTO leads ( title, description, created_at, owner_id, status)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.Exec(query, lead.Title, lead.Description, lead.CreatedAt, lead.OwnerID, lead.Status)
	return err
}

func (r *LeadRepository) Update(lead *models.Leads) error {
	query := `UPDATE leads SET title=$1, description=$2, created_at=$3, owner_id=$4, status=$5 WHERE id=$6`
	_, err := r.db.Exec(query, lead.Title, lead.Description, lead.CreatedAt, lead.OwnerID, lead.Status, lead.ID)
	return err
}

func (r *LeadRepository) GetByID(id int) (*models.Leads, error) {
	query := `SELECT id, title, description, created_at, owner_id, status FROM leads WHERE id=$1`
	row := r.db.QueryRow(query, id)
	lead := &models.Leads{}
	err := row.Scan(&lead.ID, &lead.Title, &lead.Description, &lead.CreatedAt, &lead.OwnerID, &lead.Status)
	if err != nil {
		return nil, err
	}
	return lead, nil
}
func (r *LeadRepository) Delete(id int) error {
	query := `DELETE FROM leads WHERE ID=$1`
	_, err := r.db.Exec(query, id)
	return err
}
