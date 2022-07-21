package database

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"log"
	"rmsProject/utils"
)

var RMS *sqlx.DB

type SSLMode string

const (
	SSLModeEnable  SSLMode = "enable"
	SSLModeDisable SSLMode = "disable"
)

func Connection(host, port, dbname, user, password string, sslMode SSLMode) error {

	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s  sslmode=%s", host, user, password, dbname, port, SSLModeDisable)
	log.Printf(conn)
	db, err := sqlx.Open("postgres", conn)
	utils.CheckError(err)
	err = db.Ping()
	utils.CheckError(err)
	RMS = db
	//fmt.Println("migration start")
	return migrateStart(db)
}
func migrateStart(db *sqlx.DB) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return err
	}
	m, NewErr := migrate.NewWithDatabaseInstance("file://database/migration", "postgres", driver)
	if NewErr != nil {
		return NewErr
	}
	if MigrateErr := m.Up(); MigrateErr != nil && MigrateErr != migrate.ErrNoChange { //up(): will migrate all the way up
		return err
	}
	//fmt.Println("migration successfully")
	return nil
}
