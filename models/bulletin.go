package models

import "time"

//Bulletin
type Bulletin struct {
	Id      int64
	Name    string
	Url     string
	Type    int
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func FindBulletins() ([]Bulletin, error) {
	var bulletins = make([]Bulletin, 0)
	err := orm.Asc("created").Find(&bulletins)
	return bulletins, err
}

func InsertBulletin(bulletin *Bulletin) error {
	_, err := orm.Insert(bulletin)
	return err
}

func DeleteBulletinById(id int64) error {
	_, err := orm.Id(id).Delete(new(Bulletin))
	return err
}
