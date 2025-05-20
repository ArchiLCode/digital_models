package models

type CatalogItem struct {
	UserId    string `json:"user_id,omitempty"`
	AvatarUrl string `json:"avatar_url"`
	Name      string `json:"name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}
