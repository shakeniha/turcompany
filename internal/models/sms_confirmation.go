package models

import "time"

type SMSConfirmation struct {
    ID          int64     `json:"id"`
    DocumentID  int64     `json:"document_id"`
    SMSCode     string    `json:"sms_code"`
    SentAt      time.Time `json:"sent_at"`
    Confirmed   bool      `json:"confirmed"`
    ConfirmedAt time.Time `json:"confirmed_at"`
}
