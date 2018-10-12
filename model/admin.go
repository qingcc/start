package model

import (
	"http"
)

type Admin struct {
	Id          int64  `xorm:"pk autoincr int(8)"`
	Email       string `xorm:"varchar(255)"`
	Name        string `xorm:"unique index varchar(255)"`
	Password    string `xorm:"varchar(255)"`
	AdminRoleId int64  `xorm:"index"`
	//Role          *AdminRole `xorm:"- <- ->"`
	LastLoginIp   string `xorm:"varchar(32)"`
	LastLoginTime string `xorm:"varchar(255)"`
	LoginCount    int64  `xorm:"int(8)"`
	IsLock        bool   `xorm:"default false"`
	Describe      string
	CreatedAt     app.Time `xorm:"created"`
	UpdatedAt     app.Time `xorm:"updated"`
}

type AdminList struct {
	Admin    `xorm:"extends"`
	RoleName string
	Ctime    string
}
