package repositories

import (
	"database/sql"
	"fmt"
	"turcompany/internal/models"
)

type SMSConfirmationRepository struct {
	DB *sql.DB
}

func NewSMSConfirmationRepository(db *sql.DB) *SMSConfirmationRepository {
	return &SMSConfirmationRepository{DB: db}
}

// Create вставляет новую запись SMS-подтверждения и возвращает его ID
func (r *SMSConfirmationRepository) Create(sms *models.SMSConfirmation) (int64, error) {
	query := `INSERT INTO sms_confirmations (document_id, sms_code, sent_at, confirmed, confirmed_at)
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.DB.QueryRow(query, sms.DocumentID, sms.SMSCode, sms.SentAt, sms.Confirmed, sms.ConfirmedAt).Scan(&sms.ID)
	if err != nil {
		return 0, fmt.Errorf("create sms confirmation: %w", err)
	}
	return sms.ID, nil
}

// GetByID возвращает одно подтверждение по ID
func (r *SMSConfirmationRepository) GetByID(id int64) (*models.SMSConfirmation, error) {
	query := `SELECT id, document_id, sms_code, sent_at, confirmed, confirmed_at
	          FROM sms_confirmations WHERE id = $1`
	row := r.DB.QueryRow(query, id)

	var sms models.SMSConfirmation
	err := row.Scan(&sms.ID, &sms.DocumentID, &sms.SMSCode, &sms.SentAt, &sms.Confirmed, &sms.ConfirmedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("get sms confirmation: %w", err)
	}

	return &sms, nil
}

// Update обновляет запись по ID
func (r *SMSConfirmationRepository) Update(sms *models.SMSConfirmation) error {
	query := `UPDATE sms_confirmations
	          SET document_id = $1, sms_code = $2, sent_at = $3, confirmed = $4, confirmed_at = $5
	          WHERE id = $6`
	_, err := r.DB.Exec(query, sms.DocumentID, sms.SMSCode, sms.SentAt, sms.Confirmed, sms.ConfirmedAt, sms.ID)
	if err != nil {
		return fmt.Errorf("update sms confirmation: %w", err)
	}
	return nil
}

// Delete удаляет запись по ID
func (r *SMSConfirmationRepository) Delete(id int64) error {
	query := `DELETE FROM sms_confirmations WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("delete sms confirmation: %w", err)
	}
	return nil
}
