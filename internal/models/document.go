package models

import(
  "time"
)
type DOCUMENTS struct{
  ID int64 'json:"id"'
  deal_ID int64 'json:"deal_id"'
  doc_type string 'json:"doc_type"'
  file_path string 'json:"file_path"'
  status string 'json:"status"'
  signed_at time.Time 'json:"signed_at"'
}
