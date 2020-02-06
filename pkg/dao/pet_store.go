package dao

type PetStore struct {
    // the id for this user
    //
    // required: true
    // min: 1
    ID int64 `json:"id"`

    // The name of the pet store
    //
    // required: true
    Name string `json:"name"`
}

func (e PetStore) GetID() int64 {
    return e.ID
}
