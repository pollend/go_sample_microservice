package dao

type Pet struct {
    // The id for this user
    //
    // required: true
    // min: 1
    ID int64 `json:"id"`

    // The breed of the pet
    //
    // required: true
    Breed string `json:"breed"`

    // The name of the pet
    //
    // required: true
    Name string `json:"name"`
}

func (e Pet) GetID() int64 {
    return e.ID
}
