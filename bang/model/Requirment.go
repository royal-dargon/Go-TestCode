package model

type Requirement struct {
	RequireId    string `json:"require_id" gorm:"AUTO_INCREMENT"`
	RequireTitle string `json:"require_title" gorm:"column:require_title"`
	UserId       string `json:"user_id" gorm:"column:user_id"`
	RequireNum   string `json:"require_number" gorm:"column:require_number"`
	RequireNeed  string `json:"required_need" gorm:"column:required_need"`
	RequireInfo  string `json:"require_info" gorm:"column:require_info"`
	Kind         string `json:"require_kind" gorm:"column:require_kind"`
	Time         string `json:"createtime" gorm:"column:createtime"`
}

// 获取我的发布的函数
func GetMyRequire(Id string) ([]Requirement, error) {
	var temp []Requirement
	if err := DB.Table("requires").Where("user_id = ?", Id).Find(&temp).Error; err != nil {
		return []Requirement{}, err
	}
	return temp, nil
}

// 获取我的收藏的函数
func GetMyStore(Id []string) ([]Requirement, error) {
	var temp []Requirement
	if err := DB.Table("requires").Where("require_id in (?)", Id).Find(&temp).Error; err != nil {
		return []Requirement{}, err
	}
	return temp, nil
}

// 通过id获得招募信息
func GetRequireInfo(Id string) (Requirement, error) {
	var temp Requirement
	if err := DB.Table("requires").Where("require_id = ?", Id).Find(&temp).Error; err != nil {
		return Requirement{}, err
	}
	return temp, nil
}
