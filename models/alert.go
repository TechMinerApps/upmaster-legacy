package models

// swagger:parameters Alert
type Alert struct {
	BaseModel
	AlertChannal   *AlertChannal `json:"alert_channal,omitempty"`
	AlertChannalID uint          `json:"alert_channal_id"`
	Status         uint          `json:"status"`
	User           *User         `json:"user,omitempty"`
	UserID         uint          `json:"user_id"`
}

// swagger:parameters AlertChannal
type AlertChannal struct {
	BaseModel
	Name      string  `json:"name"`
	Alerts    []Alert `json:"alerts"`
	User      *User   `json:"user,omitempty"`
	UserID    uint    `json:"user_id"`
	IsEnabled bool    `json:"is_enabled"`
	Type      uint    `json:"type"`
	Config    []byte  `json:"config"`
}
