package services

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

var (
	insertQuery = "INSERT INTO `booleans` (`name`,`val`,`uuid`) VALUES (?,?,?)"
	fetchQuery  = "SELECT * FROM `booleans`  WHERE (uuid = ?) LIMIT 1"
	updateQuery = "UPDATE `booleans` SET `key` = ?, `uuid` = ?, `value` = ? WHERE (uuid = ?)"
	deleteQuery = "DELETE FROM `booleans` WHERE (uuid = ? )"
)

func TestAddToDB(t *testing.T) {
	d, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	db, err = gorm.Open("mysql", d)
	if err != nil {
		log.Fatal("Unable to initialize the datbase setup")
	}
	repoImpl := RepoImpl{}
	MyRepo = repoImpl

	rows := sqlmock.
		NewRows([]string{"key", "value", "id"}).
		AddRow("ash", true, "68df2cbb-a432-4b35-8a99-2cf3de9b243c")

	mock.ExpectBegin()
	mock.ExpectExec(insertQuery).WithArgs("ash", true, sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()
	mock.ExpectQuery(fetchQuery).WithArgs(sqlmock.AnyArg()).WillReturnRows(rows)

	MyRepo.AddToDB("ash", true)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateInDB(t *testing.T) {
	d, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	db, err = gorm.Open("mysql", d)
	if err != nil {
		log.Fatal("Unable to initialize the datbase setup")
	}
	repoImpl := RepoImpl{}
	MyRepo = repoImpl

	rows := sqlmock.
		NewRows([]string{"key", "value", "id"}).
		AddRow("ash", true, "68df2cbb-a432-4b35-8a99-2cf3de9b243c")

	mock.ExpectBegin()
	mock.ExpectExec(updateQuery).WithArgs("ash", sqlmock.AnyArg(), true, sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()
	mock.ExpectQuery(fetchQuery).WithArgs(sqlmock.AnyArg()).WillReturnRows(rows)

	if _, err = MyRepo.UpdateInDB("ash", true, "68df2cbb-a432-4b35-8a99-2cf3de9b243c"); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteFromDB(t *testing.T) {
	d, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	db, err = gorm.Open("mysql", d)
	if err != nil {
		log.Fatal("Unable to initialize the datbase setup")
	}
	repoImpl := RepoImpl{}
	MyRepo = repoImpl

	mock.ExpectBegin()
	mock.ExpectExec(deleteQuery).WithArgs(sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()
	if err = MyRepo.DeleteFromDB("68df2cbb-a432-4b35-8a99-2cf3de9b243c"); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
