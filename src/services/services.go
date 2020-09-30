package services

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"

	//This import is for proper functioning of gorm with mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Repo is something
type Repo interface {
	AddToDB(name string, value bool) Boolean
	GetFromDB(uuid string) (Boolean, error)
	UpdateInDB(name string, value bool, uuid string) (Boolean, error)
	DeleteFromDB(uuid string) error
}

// RepoImpl is nothing
type RepoImpl struct{}

// Boolean is the model for our database schema
type Boolean struct {
	gorm.Model
	Name string
	Val  bool
	UUID string `gorm:"primary_key"`
}

// MyRepo is a variable
var MyRepo Repo = RepoImpl{}
var db *gorm.DB

const (
	user     = ""
	password = ""
	table    = ""
)

// Init will setup the database and create table
func Init() {
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", user, password, table))
	if err != nil {
		log.Fatal(err)
	}
	db.Table("booleans").AutoMigrate(&Boolean{})
}

// AddToDB is for making an entry into the database
func (repo RepoImpl) AddToDB(name string, value bool) Boolean {
	b := Boolean{
		UUID: uuid.New().String(),
		Val:  value,
		Name: name,
	}
	db.Table("booleans").Create(&b)
	return b
}

// GetFromDB is to get an entry from the database
func (repo RepoImpl) GetFromDB(uuid string) (Boolean, error) {
	var b Boolean
	if res := db.Table("booleans").Where("uuid = ?", uuid).First(&b); res.Error != nil {
		return Boolean{}, res.Error
	}
	return b, nil
}

// UpdateInDB does what its name suggests
func (repo RepoImpl) UpdateInDB(name string, value bool, uuid string) (Boolean, error) {
	var b Boolean
	res := db.Table("booleans").Where("uuid = ?", uuid).First(&b)
	if res.Error != nil {
		return Boolean{}, res.Error
	}
	if name != "" {
		b.Name = name
	}
	b.Val = value
	db.Save(&b)
	return b, nil
}

// DeleteFromDB sets the deleted_at column to the current timestamp, thus making the entry unavailable
func (repo RepoImpl) DeleteFromDB(uuid string) error {
	res := db.Table("booleans").Where("uuid = ?", uuid)
	if res.Error == nil {
		res.Delete(&Boolean{})
	}
	return res.Error
}
