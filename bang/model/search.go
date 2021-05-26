package model

import "fmt"

type SearchContent struct {
	Content string `json:"content"`
}

// 搜索的函数
func SearchOK(content string) (Contest, error) {
	var temp Contest
	if err := DB.Table("contest").Where("contest_name = ?", content).Find(&temp).Error; err != nil {
		fmt.Println(err)
		return Contest{}, err
	}
	return temp, nil
}

// 搜索招募信息的函数
func GetResult2(content string) ([]Requirement, error) {
	var temp []Requirement
	if err := DB.Table("requirements").Where("requiremnets = ?", content).Limit(5).Find(&temp).Error; err != nil {
		return nil, err
	}
	return temp, nil
}
