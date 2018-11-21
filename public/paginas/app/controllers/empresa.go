package controllers

import (
	"github.com/revel/revel"

	"paginas/app/routes"
	//"paginas/app/models"
	"paginas/app"
	//"fmt"
)

type Empresa struct {
	*revel.Controller
	Id int
	Name string
	Rua string
	Cidade string
	Pais string
	Longitude string 
	Latitude string
}

func (c Empresa) Index() revel.Result {

	return c.Render()
}

func (c Empresa) InsertView() revel.Result {
	
	sqlStatement := `SELECT id, nome, rua, cidade, pais, longitude, latitude FROM empresa`
	
	rows, err := app.DB.Query(sqlStatement)
	
	if err != nil {
    }
    defer rows.Close()

    var result []Empresa
    for rows.Next() {
		var id int
		var nome, rua, cidade, pais, longitude, latitude string
        err := rows.Scan(&id, &nome, &rua, &cidade, &pais, &longitude, &latitude)
        if err != nil {
            
        }

        result = append(result, Empresa{Id: id, Name: nome, Rua: rua, Cidade: cidade, Longitude: longitude, Latitude: latitude})
    }
    if err := rows.Err(); err != nil {
      
    }
	
	return c.Render(result)
		
}

func (c Empresa) Insert(name, address, city, country, longitude, latitude string) revel.Result {
	
	sql := "INSERT INTO empresa (nome, rua, cidade, pais, longitude, latitude) VALUES ($1,$2,$3,$4,$5,$6)"
	err := app.DB.QueryRow(sql,name, address, city, country, longitude, latitude)
	
	if err == nil {
		panic(err)
	}
	


	c.Flash.Success("Insert success")
	return c.Redirect(routes.Empresa.InsertView())
	
}

func (c Empresa) Delete(id int) revel.Result {
	sqlStatement := `DELETE FROM empresa WHERE id = $1;`
	_, err := app.DB.Exec(sqlStatement, id)
	if err != nil {
	panic(err)
	}

	return c.Redirect(routes.Empresa.Index())
}

func (c Empresa) Show(nome, cidade string) revel.Result {
	sqlStatement := "SELECT id, nome, rua, cidade, pais, longitude, latitude FROM empresa WHERE nome LIKE '%' || $1 || '%' and cidade LIKE '%' || $2 || '%'"
	var name, address,city,country,longitude,latitude string
	var id1 int
	row := app.DB.QueryRow(sqlStatement, nome, cidade)
	err := row.Scan(&id1, &name, &address, &city, &country, &longitude,&latitude);
	if err != nil {
		
	}
	return c.Render(id1,name, address,city,country,longitude,latitude)
}

func (c Empresa) List() revel.Result {
	sqlStatement := `SELECT id, nome, rua, cidade, pais, longitude, latitude FROM empresa`
	
	rows, err := app.DB.Query(sqlStatement)
	if err != nil {
		// handle this error better than this
		panic(err)
	  }
	
	var empresas [][]string
	cols, _ := rows.Columns()
	pointers := make([]interface{}, len(cols))
	container := make([]string, len(cols))
	for i, _ := range pointers {
		pointers[i] = &container[i]
	}
	for rows.Next() {
		rows.Scan(pointers...)
		empresas = append(empresas, container)
	}

	return c.Render(empresas)
}
