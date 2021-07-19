package model

type AuthorityMenu struct {
	Menu
	MenuID      string          `json:"menuId" gorm:"comment:菜单ID"`
	AuthorityId string          `json:"-" gorm:"comment:角色ID"`
	Children    []AuthorityMenu `json:"children" gorm:"-"`
	Parameters  []MenuParameter `json:"parameters" gorm:"foreignKey:MenuID;references:MenuID"`
}

func (s *AuthorityMenu) TableName() string {
	return "system_authority_menu"
}
