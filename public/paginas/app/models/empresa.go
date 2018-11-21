package models

import (
	"github.com/revel/revel"
)

type Empresa struct {
	EmpresaId          int
	Name, Address    string
	City			 string
	Country          string
	Longitude, Latitude string
}

func (empresa *Empresa) Validate(v *revel.Validation) {
	v.Check(empresa.Name,
		revel.Required{},
		revel.MaxSize{50},
	)

	v.MaxSize(empresa.Address, 100)

	v.Check(empresa.City,
		revel.Required{},
		revel.MaxSize{40},
	)

	v.Check(empresa.Country,
		revel.Required{},
		revel.MaxSize{40},
		revel.MinSize{2},
	)

	v.Check(empresa.Longitude,
		revel.Required{},
		revel.MaxSize{40},
		revel.MinSize{2},
	)

	v.Check(empresa.Latitude,
		revel.Required{},
		revel.MaxSize{40},
		revel.MinSize{2},
	)
}
