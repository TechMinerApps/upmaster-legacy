package models

// swagger:response Endpoint
type Endpoint struct {
	BaseModel
	// in:body
	Name      string  `json:"name"`
	User      *User   `json:"user,omitempty"`
	UserID    uint    `json:"user_id"`
	URL       string  `json:"url"`
	Interval  uint    `json:"interval"`
	IsEnabled bool    `json:"is_enabled"`
	IsPublic  bool    `json:"is_public"`
	Alerts    []Alert `json:"alerts"`
}
