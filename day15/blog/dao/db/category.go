package db

import "github.com/sealandsigh/gotest5/day15/blog/model"

// 添加分类

func InsertCategory(category *model.Category) (categoryId int64, err error) {
	sqlstr := "insert into category(category_name,category_no) value(?,?)"
	result, err := DB.Exec(sqlstr, category.CategoryName, category.CategoryNo)
	if err != nil {
		return
	}
	categoryId, err = result.LastInsertId()
	return
}

// 获取单个分类

func GetCategoryById(id int64) (category *model.Category, err error) {
	category = &model.Category{}
	sqlstr := "select id,category_name,category_no from category where id=?"
	err = DB.Get(category, sqlstr, id)
	return
}
