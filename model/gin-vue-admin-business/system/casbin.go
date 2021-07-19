package model

type Casbin struct {
	Path        string `gorm:"column:v1"`
	PType       string `gorm:"column:ptype"`
	Method      string `gorm:"column:v2"`
	AuthorityId string `gorm:"column:v0"`
}

func (c *Casbin) TableName() string {
	return "casbin_rule"
}
