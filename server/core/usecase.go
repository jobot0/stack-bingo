package core

import "fmt"

type createDumbUsecase struct {
	dumbRepository DumbRepository 
}

type CreateDumbUsecase interface {
	Execute(*Dumb) error
}

func NewCreateDumbUsecase(d DumbRepository) CreateDumbUsecase{
	return &createDumbUsecase{
		dumbRepository: d,
	}
}

func (c *createDumbUsecase) Execute(dumb *Dumb) (err error) {
	fmt.Printf("%+v\n", dumb)	
	return 
}
