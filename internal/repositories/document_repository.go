package repositories

import (
    "database/sql"
    "fmt"
    "turcompany/internal/models"
)

type DocumentRepository struct {
    DB *sql.DB
}

func NewDocumentRepository(db *sql.DB) *DocumentRepository {
    return &DocumentRepository{DB: db}
}

func (r *DocumentRepository) Create(doc *models.Document) (int64, error) {
    query := `INSERT INTO documents (deal_id, doc_type, file_path, status, signed_at)
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
    err := r.DB.QueryRow(query, doc.DealID, doc.DocType, doc.FilePath, doc.Status, doc.SignedAt).Scan(&doc.ID)
    if err != nil {
        return 0, fmt.Errorf("create document: %w", err)
    }
    return doc.ID, nil
}

func (r *DocumentRepository) GetByID(id int64) (*models.Document, error) {
    query := `SELECT id, deal_id, doc_type, file_path, status, signed_at FROM documents WHERE id = $1`
    row := r.DB.QueryRow(query, id)
    var doc models.Document
    err := row.Scan(&doc.ID, &doc.DealID, &doc.DocType, &doc.FilePath, &doc.Status, &doc.SignedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, fmt.Errorf("get document: %w", err)
    }
    return &doc, nil
}

func (r *DocumentRepository) Update(doc *models.Document) error {
    query := `UPDATE documents SET deal_id=$1, doc_type=$2, file_path=$3, status=$4, signed_at=$5 WHERE id=$6`
    _, err := r.DB.Exec(query, doc.DealID, doc.DocType, doc.FilePath, doc.Status, doc.SignedAt, doc.ID)
    if err != nil {
        return fmt.Errorf("update document: %w", err)
    }
    return nil
}

func (r *DocumentRepository) Delete(id int64) error {
    query := `DELETE FROM documents WHERE id = $1`
    _, err := r.DB.Exec(query, id)
    if err != nil {
        return fmt.Errorf("delete document: %w", err)
    }
    return nil
}
