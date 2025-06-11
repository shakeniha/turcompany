package repositories

import (
	"database/sql"
	"turcompany/internal/models"
)

type DealRepository struct {
	db *sql.DB
}

func  NewDealRepository(db *sql.DB) *DealRepository {
	return &DealRepository{db :db}
}
func (r *DealRepository) Create(deal *models.Deals) error{
	query := `INSERT INTO deals (id, lead_id, amount, currency, status, created_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, deal.ID, deal.LeadID, &deal.Amount, deal.Currency, deal.Status, &deal.CreatedAt)
	return err
}
func (r *DealRepository) Update(deal *models.Deals) error {
	query :=  `UPDATE deals SET lead_id=$1, amount=$2, currency=$3, status=$4 WHERE id=$5`
	_, err := r.db.Exec(query, deal.LeadID, &deal.Amount, deal.Currency, deal.Status, deal.ID)
	return err
}

func (r *DealRepository) GetByID(id string)(*models.Deals, error) {
	query:= `SELECT id, lead_id, amount, currency, status, created_at FROM deals WHERE id=$1`
	row := r.db.QueryRow(query, id)
	deal:=&models.Deals{}
	err:= row.Scan(&deal.ID, &deal.LeadID, &deal.Amount, &deal.Currency, &deal.Status, &deal.CreatedAt)
	if err != nil {
		return nil,err
	}
	return deal, nil
}
func (r *DealRepository)Delete(id string) error {
	query := `DELETE FROM deals WHERE id=$1`
	_, err :=r.db.Exec(query,id)
	return err
}