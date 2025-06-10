package models
import(
  "time"
)

type SMS_CONFIRMATIONS struct{
  ID int64 'json:"id"'
  document_id int64 'json:"document_id"'
  sms_code string 'json:"sms_code"'
  sent_at time.Time 'json:"sent_at"'
  confirmed bool 'json:"confirmed"'
  confirmed_at time.Time 'json:"confirmed_at"'
}
