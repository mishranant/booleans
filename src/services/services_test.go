package services

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func TestAddToDB(t *testing.T) {
	// configure to use case sensitive SQL query matcher
	// instead of default regular expression matcher
	d, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	db, err = gorm.Open("mysql", d)
	db.Debug()
	if err != nil {
		log.Fatal("Unable to initialize the datbase setup")
	}

	rows := sqlmock.
		NewRows([]string{"key", "value", "id"}).
		AddRow("ash", true, "68df2cbb-a432-4b35-8a99-2cf3de9b243c")

	// Update
	mock.ExpectBegin()
	mock.
		ExpectExec("INSERT INTO `booleans` (`name`,`val`,`uuid`) VALUES (?,?,?)").
		WithArgs("ash", true, sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	// Fetch
	mock.
		ExpectQuery("SELECT * FROM `booleans` WHERE (uuid = ?) ORDER BY `booleans`.`uuid` ASC LIMIT 1").
		WithArgs(sqlmock.AnyArg()).WillReturnRows(rows)

	MyRepo.AddToDB("ash", true)
	MyRepo.GetFromDB("some")

	// we make sure that all expectations were met
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateInDB(t *testing.T) {
	d, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	db, err = gorm.Open("mysql", d)
	if err != nil {
		log.Fatal("Unable to initialize the datbase setup")
	}

	rows := sqlmock.
		NewRows([]string{"key", "value", "id"}).
		AddRow("ash", true, "68df2cbb-a432-4b35-8a99-2cf3de9b243c")

	// Update
	mock.ExpectBegin()
	mock.
		ExpectExec("UPDATE `booleans` SET `name` = ?, `val` = ? WHERE (uuid = ?)").
		WithArgs("bash", false, sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()

	// Fetch
	mock.
		ExpectQuery("SELECT * FROM `booleans` WHERE (uuid = ?) ORDER BY `booleans`.`uuid` ASC LIMIT 1").
		WithArgs(sqlmock.AnyArg()).WillReturnRows(rows)

	if _, err = MyRepo.UpdateInDB("bash", false, "68df2cbb-a432-4b35-8a99-2cf3de9b243c"); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteFromDB(t *testing.T) {
	d, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	db, err = gorm.Open("mysql", d)
	if err != nil {
		log.Fatal("Unable to initialize the datbase setup")
	}

	mock.ExpectBegin()
	mock.
		ExpectExec("DELETE FROM `booleans` WHERE (uuid = ?)").
		WithArgs(sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()

	if err = MyRepo.DeleteFromDB("68df2cbb-a432-4b35-8a99-2cf3de9b243c"); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
