package controllers

import (
	"github.com/revel/revel"

	"paginas/app/routes"
	//"paginas/app/models"
	//"paginas/app"
	//"fmt"
)

type Login struct {
	*revel.Controller
	
}

func (c Login) Index() revel.Result {

	return c.Render()
}

func (c Login) Auth(username, password string) revel.Result {
	
	if (username == "admin") && (password=="admin") {
			c.Session["user"] = username
			c.Session.SetNoExpiration()
			c.Flash.Success("Welcome, " + username)
			return c.Redirect(routes.Empresa.InsertView())
	}

	c.Flash.Out["username"] = username
	c.Flash.Error("Login failed")
	return c.Redirect(routes.Login.Index())
}