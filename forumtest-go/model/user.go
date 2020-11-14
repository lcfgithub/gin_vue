package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name" gorm:"varchar(20) not null"`
	Telphone string `json:"telphone" gorm:"char(11) not null unique"`
	Password string `json:"password" gorm:"size:255 not null"`
}
