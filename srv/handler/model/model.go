package model

import (
	__ "guanli/proto"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(30)"`
	Password string `gorm:"type:varchar(32)"`
}

func (u *User) FindUser(db *gorm.DB, name string) error {
	return db.Debug().Where("name = ?", name).Find(&u).Error
}

func (u *User) UserAdd(db *gorm.DB) error {
	return db.Debug().Create(&u).Error
}

type Role struct {
	gorm.Model
	Uid  int    `gorm:"type:int"`
	Name string `gorm:"type:varchar(30)"`
}

func (r *Role) FindRole(db *gorm.DB, name string) error {
	return db.Debug().Where("name = ?", name).Find(&r).Error
}

func (r *Role) RoleAdd(db *gorm.DB) error {
	return db.Debug().Create(&r).Error
}

func (r *Role) RoleList(db *gorm.DB, id int32) (error, []*__.RoleList) {
	var list []*__.RoleList
	tx := db.Debug().Where("id = ?", id).Joins("users.*").Joins("roles.*").
		Find(r).Error
	return tx, list
}

func (r *Role) FindName(db *gorm.DB, id int32) error {
	return db.Debug().Where("id = ?", id).Find(&r).Error
}

func (r *Role) DeleteRole(db *gorm.DB, id int32) error {
	return db.Debug().Where("id = ?", id).Delete(&r).Error
}

type QuanXian struct {
	gorm.Model
	Rid  int    `gorm:"type:int"`
	Uid  int    `gorm:"type:int"`
	Name string `gorm:"type:varchar(30)"`
}

func (x *QuanXian) FindQuan(db *gorm.DB, name string) error {
	return db.Debug().Where("name = ?", name).Find(&x).Error
}

func (x *QuanXian) QuanAdd(db *gorm.DB) error {
	return db.Debug().Create(&x).Error
}

type User_Role struct {
	gorm.Model
	Rid int `gorm:"type:int"`
	Uid int `gorm:"type:int"`
}
type Role_QuanXian struct {
	gorm.Model
	Rid int `gorm:"type:int"`
	Qid int `gorm:"type:int"`
}
