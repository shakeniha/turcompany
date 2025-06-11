package repositories
import (
	"database/sql"
	"turcompany/internal/models"
)

type LeadRepository struct {
	db *sql.DB
}

func NewLeadRepository(db *sql.DB)* LeadRepository {
	return &LeadRepository{db:db}
}
func (r *LeadRepository) Create(lead *models.Leads)error{
	query:= `INSERT INTO leads (ID,Title,Descriptin, CreateAt, OwnerID, Status,) VALUES ($1, $2, $3, $4, $5, $6)`
	_,err :=r.db.Exec(query, lead.ID, lead.Title, lead.Description, &lead.CreatedAt, lead.OwnerID, lead.Status)
	return err
}
func (r *LeadRepository) Update(lead *models.Leads) error {
	query := `UPDATE leads SET Title=$1, Description=$2, CreatedAt=$3, OwnerID=$4, Status=$5 WHERE ID=$6`
	_, err := r.db.Exec(query, lead.Title, lead.Description, &lead.CreatedAt, lead.OwnerID, lead.Status, lead.ID)
	return err
}

func (r *LeadRepository) GetByID(id string) (*models.Leads, error) {
	query := `SELECT ID, Title, Description, CreatedAt, OwnerID, Status FROM leads WHERE ID=$1`
	row := r.db.QueryRow(query, id)
	lead := &models.Leads{}
	err := row.Scan(&lead.ID, &lead.Title, &lead.Description, &lead.CreatedAt, &lead.OwnerID, &lead.Status)
	if err != nil {
		return nil, err
	}
	return lead, nil
}
func (r *LeadRepository) Delete(id string) error {
	query := `DELETE FROM leads WHERE ID=$1`
	_, err := r.db.Exec(query, id)
	return err
}

