package model

type UserAndId struct {
	Id        string `json:"id" gorm:"id"`
	User_id   string `json:"user_id" gorm:"column:user_id"`
	RequireId string `json:"require_id" gorm:"column:require_id"`
}

func GetMyStoreId(Id string) ([]string, error) {
	var UserId []string
	var temp []UserAndId
	if err := DB.Table("user_require").Where("user_id = ?", Id).Find(&temp).Error; err != nil {
		return nil, err
	}
	for _, id := range temp {
		UserId = append(UserId, string(id.RequireId))
	}
	return UserId, nil
}
