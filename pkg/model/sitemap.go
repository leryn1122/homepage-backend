package model

type Sitemap struct {
	CommonModel
	Name        string `json:"name" gorm:"column:name"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	URL         string `json:"url" gorm:"column:url"`
	Logo        string `json:"logo" gorm:"column:logo"`
	Enabeld     bool   `json:"enabled" gorm:"column:enabled"`
}

func (Sitemap) TableName() string {
	return "sitemap"
}
