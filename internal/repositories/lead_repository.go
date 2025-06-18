package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"turcompany/internal/models"
)

type LeadRepository struct {
	db *sql.DB
}

func NewLeadRepository(db *sql.DB) *LeadRepository {
	if db == nil {
		log.Fatalf("Received nil database connection")
	}
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
func (r *LeadRepository) CountLeads() (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM leads"
	err := r.db.QueryRow(query).Scan(&count)
	return count, err
}
func (r *LeadRepository) FilterLeads(status string, ownerID string) ([]models.Leads, error) {
	query := "SELECT id, title, description, created_at, owner_id, status FROM leads WHERE 1=1"
	args := []interface{}{}
	i := 1

	if status != "" {
		query += fmt.Sprintf(" AND status = $%d", i)
		args = append(args, status)
		i++
	}

	if ownerID != "" {
		query += fmt.Sprintf(" AND owner_id = $%d", i)
		args = append(args, ownerID)
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var leads []models.Leads
	for rows.Next() {
		var lead models.Leads
		if err := rows.Scan(&lead.ID, &lead.Title, &lead.Description, &lead.CreatedAt, &lead.OwnerID, &lead.Status); err != nil {
			return nil, err
		}
		leads = append(leads, lead)
	}
	return leads, nil
}
