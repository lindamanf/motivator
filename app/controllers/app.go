package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	// models.AddUser(&User{Name: "lindamanf", mail: "lindaman.h.12@outlook.jp", password: "P@ssw0rd"})
	return c.Render()
}
