package repositories

import (
    "github.com/jmoiron/sqlx"
    "prospect/mobile-physician-api/pkg/dao"
)

type PetStoreInterface interface {
    Add(tx* sqlx.Tx,entity* dao.PetStore)
    Delete(tx* sqlx.Tx, entity dao.PetStore)
    Update(entity dao.PetStore)
}

type PetStoreRepository struct {
}


func NewPetStoreRepository() PetStoreInterface {
    return &PetStoreRepository{}
}

func (p PetStoreRepository) GetProviders() []dao.PetStore  {
    //tx := p.db.MustBegin()
    //tx.MustExec("SELECT * FROM provider");
    return nil
}

func (p PetStoreRepository) GetAllProviders() []dao.PetStore {
    return nil
}

func (p PetStoreRepository) Add(tx* sqlx.Tx,entity* dao.PetStore) {
    tx.MustExec("INSERT into provider () VALUE ()", entity.Name)
}

func (p PetStoreRepository) Delete(tx* sqlx.Tx, entity dao.PetStore){

}

func (p PetStoreRepository) Update(entity dao.PetStore) {
}
