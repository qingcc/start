package model

import (
	"http"
)

type AdminRole struct {
	Id        int64    `xorm:"pk autoincr int(8)"`
	RoleName  string   `xorm:"unique varchar(255)"`
	IsSuper   bool     `xorm:"default false"`
	IsSys     bool     `xorm:"default true"`
	Describe  string   `xorm:"varchar(255)"`
	CreatedAt app.Time `xorm:"created"`
	UpdatedAt app.Time `xorm:"updated"`
}
