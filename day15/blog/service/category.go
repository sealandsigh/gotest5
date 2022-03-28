package service

import (
	"github.com/sealandsigh/gotest5/day15/blog/dao/db"
	"github.com/sealandsigh/gotest5/day15/blog/model"
)

// 获取所有分类

func GetALLCategoryList() (categoryList []*model.Category, err error) {
	categoryList, err = db.GetAllCategoryList()
	if err != nil {
		return
	}
	return
}
