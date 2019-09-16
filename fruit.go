package main

type Fruit struct {
	Id     int64             `json:"id"`
	Colors map[string]string `json:"colors"`
	Stores map[string]*Store `json:"stores"`
}

type Store struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func (d *Fruit) Create() (int64, error) {
	return Db.Insert(d)
}

func (Fruit) GetById(id int64) (bool, Fruit, error) {
	fruit := Fruit{}
	has, err := Db.Where("id=?", id).Get(&fruit)
	return has, fruit, err
}
