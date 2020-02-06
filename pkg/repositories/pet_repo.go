package repositories

import (
    "github.com/jmoiron/sqlx"
    "prospect/mobile-physician-api/pkg/dao"
)

type PetInterface interface {
    Add(tx* sqlx.Tx,entity* dao.Pet)
    Delete(tx* sqlx.Tx, entity dao.Pet)
    Update(entity dao.Pet)
}

type PetRepository struct {
}

func NewPetRepository() PetInterface {
    return &PetRepository{}
}

func (p PetRepository) GetProviders() []dao.Pet  {
    //tx := p.db.MustBegin()
    //tx.MustExec("SELECT * FROM provider");
    return nil
}

func (p PetRepository) GetAllProviders() []dao.Pet {
    return nil
}

func (p PetRepository) Add(tx* sqlx.Tx,entity* dao.Pet) {
    tx.MustExec("INSERT into provider () VALUES ()", entity.Breed, entity.Name)
}

func (p PetRepository) Delete(tx* sqlx.Tx, entity dao.Pet){

}

func (p PetRepository) Update(entity dao.Pet) {
}
