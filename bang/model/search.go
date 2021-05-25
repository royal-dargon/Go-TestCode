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
