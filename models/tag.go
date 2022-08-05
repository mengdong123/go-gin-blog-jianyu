package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	DeletedOn  string `json:"deleted_on"`
	State      int    `json:"state"`
}

// CleanAllTag cron用到的定时任务清楚tag的state不是0的数据
func CleanAllTag() bool {
	// 注意硬删除要使用 Unscoped()，这是 GORM 的约定
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tag{})

	return true
}

// DeleteTagById 根据id删除tag
func DeleteTagById(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}

// EditTag 更新tag
func EditTag(id int, data map[string]interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	return true
}

// ExistTagById 根据id判断tag是否存在
func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

// 设置created_on的值
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

// 设置momodified_on的值
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	return true
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}
