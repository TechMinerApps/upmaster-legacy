package models

type User struct {
	BaseModel

	Username      string         `json:"username" gorm:"index"`
	Password      string         `json:"password,omitempty"`
	Email         string         `json:"email"`
	IsAdmin       bool           `json:"is_admin"`
	Endpoints     []Endpoint     `json:"endpoints"`
	Alerts        []Alert        `json:"alerts"`
	AlertChannals []AlertChannal `json:"alert_channals"`
}
