package controllers

import (
	"github.com/revel/revel"
	"paginas/app"
	"fmt"
)

type App struct {
	*revel.Controller
	
}

func (c App) Index() revel.Result {
	//vart := "Aloha World"

	fmt.Println("ROWS1:")
	sql := "SELECT id, username from userstb where id=0 "
	rows, err := app.DB.Query(sql)
	
	defer rows.Close()
  	for rows.Next() {
		var id int
		var Name string
		err = rows.Scan(&id, &Name)
		if err != nil {
			// handle this error
			panic(err)
		}
		
		fmt.Println(id, Name)
  	}

	return c.Render(rows)
}
