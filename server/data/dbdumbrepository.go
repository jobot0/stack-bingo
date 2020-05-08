package data

import (
	"github.com/go-pg/pg/v9"
	"stack-bingo/core"
)

type Dumbs struct {
	tableName struct{} `pg:"dumbs"`
	Name	string
}

type dbDumbRepository struct {
	DbConn *pg.DB
}

func NewDbDumbRepository(db *pg.DB) core.DumbRepository {
	return &dbDumbRepository{
		DbConn: db,
	}
}

func (ddr *dbDumbRepository) Save(dumb *core.Dumb) (err error) {
	d := Dumbs {
		Name: dumb.Name,
	}

	err = ddr.DbConn.Insert(&d)
	if err != nil {
		panic(err)
	}
	return 
}
