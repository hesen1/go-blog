package user

import (
	"github.com/jinzhu/gorm"

	"github.com/hesen/blog/pool"
)

// User 用户表
type User struct {
	gorm.Model
	Name         string `form:"name" binding:"required"`
  Age          *int `gorm:"type:int" form:"age"`
  Email        string  `gorm:"type:varchar(100);unique_index" form:"email" binding:"required,email"`
  Role         string  `gorm:"size:5;default:'user'" form:"role"` // set field size to 255
  Phone string `gorm:"unique;not null" form:"phone" binding:"required,phone"` // set member number to unique and not null
  Address      string  `gorm:"type:varchar(100)" form:"address"` // create index with name `addr` for address
	Password string `gorm:"type:varchar(40);not null" form:"password" binding:"required"`
	IP         string `gorm:"type:varchar(40)"`
}

// NewUserModel 新建一个User实列
func NewUserModel() User {
	return User{}
}

func (user *User) genSelectFields() []string {
	return []string{ "name", "age", "email", "role", "phone", "address", "ip" }
}

func (user *User) genBaseWheres() User {
	where := NewUserModel()

	if user.Phone != "" {
		where.Phone = user.Phone
	}

	return where
}

// 查找用户
func (user *User) findUsers(page int, limit int) {
	pool.DB.Find(user).Limit(limit).Offset((page - 1) * limit)
}

// 获取用户信息，不包含密码
func (user *User) findUser(out interface{}) error {
	db := pool.DB.Select(user.genSelectFields()).Where(user.genBaseWheres())

	if user.Email != "" {
		db = db.Or("email = ?", user.Email)
	}
	return pool.FliterError(db.First(out).Error)
}

// 包含密码
func (user *User) findUserWithFullInfo(out interface{}) error {
	db := pool.DB.Where(user.genBaseWheres())

	if user.Email != "" {
		db = db.Or("email = ?", user.Email)
	}
	return pool.FliterError(db.First(out).Error)
}

func (user *User) create() error {
	return pool.DB.Create(user).Error
}
