package models
import "google.golang.org/genproto/googleapis/type/date"

type Leads struct {
    ID          string              `json:"id"`
    Title       string              `json:"title"`
    Description string              `json:"description"`
    CreatedAt   date.Date           `json:"created_at"`
    OwnerID     int                 `json:"owner_id"`
    Status      string              `json:"status"`
}