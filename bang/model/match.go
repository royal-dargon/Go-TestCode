package model

import "fmt"

type Contest struct {
	ContestId    string `json:"contest_id" gorm:"AUTO_INCREMENT"`
	ConetestName string `json:"contest_name" gorm:"column:contest_name"`
	ContestInfo  string `json:"contest_info" gorm:"column:contest_info"`
	ContestNum   string `json:"contest_people" gorm:"column:contest_people"`
	CreateTime   string `json:"createtime" gorm:"column:createtime"`
	ContestUrl   string `json:"contest_url" gorm:"column:contest_url"`
	ContestKind  string `json:"contest_kind" gorm:"column:contest_kind"`
}

func GetMatchInfo() ([]Contest, error) {
	var temp []Contest
	if err := DB.Table("contest").Limit(10).Find(&temp).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return temp, nil
}

// 获取单个比赛详细信息的函数
func GetContestInfo(id string) (Contest, error) {
	var temp Contest
	if err := DB.Table("contest").Where("contest_id = ?", id).Find(&temp).Error; err != nil {
		fmt.Println(err)
		return Contest{}, err
	}
	return temp, nil
}
