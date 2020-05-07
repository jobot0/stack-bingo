package data

import (
	"fmt"
	"github.com/go-pg/pg/v9"
)

type Dumbs struct {
	tableName struct{} `pg:"dumbs"`
	Id	int64 
	Name	string
}

func (d Dumbs) String() string {
	return fmt.Sprintf("Dumb<%d %s>", d.Id, d.Name)
}

func DB_Model() {
	db := pg.Connect(&pg.Options{
		Database: "stack-bingo",
		User: "postgres",
		Password: "Pass2020!",
	})
	defer db.Close()

	dumb := Dumbs{
		Id: 1,
		Name: "toto",
	}
	err := db.Select(&dumb)
	if err != nil {
		panic(err)
	}

	fmt.Println(dumb)
}

