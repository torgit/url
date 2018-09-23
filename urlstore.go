package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type url struct {
	id          int
	originalUrl string
	shortUrl    string
}

type urlStore interface {
	getById(int) (url, error)
	getByOriginalUrl(string) (url, error)
	upsertUrl(url) error
}

type dbStore struct {
	db                  *sql.DB
	getByIdStm          *sql.Stmt
	getByOriginalUrlStm *sql.Stmt
	upsertUrlStm        *sql.Stmt
}

var errorInitializing = fmt.Errorf("initailize db error")
var errorConnection = fmt.Errorf("get db connection error")
var errorFailToFind = fmt.Errorf("failed to find url")
var errorFailToSave = fmt.Errorf("failed to save url")

func newDbStore() (*dbStore, error) {
	mysqlIp := os.Getenv("mysqlIp")
	mysqlPort := os.Getenv("mysqlPort")
	mysqlUser := os.Getenv("mysqlUser")
	mysqlPassword := os.Getenv("mysqlPassword")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/taskworld", mysqlUser, mysqlPassword, mysqlIp, mysqlPort)

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		return &dbStore{}, errorInitializing
	}
	err = db.Ping()
	if err != nil {
		return &dbStore{}, errorConnection
	}

	getByIdStm, err1 := db.Prepare("SELECT id, originalUrl, shortUrl FROM url WHERE id = ?")
	getByOriginalUrlStm, err2 := db.Prepare("SELECT id, originalUrl, shortUrl FROM url WHERE originalUrl = ?")
	upsertUrlStm, err3 := db.Prepare("INSERT INTO url (id, originalUrl, shortUrl) values (?, ?, ?) ON DUPLICATE KEY UPDATE shortUrl = ?")

	if err1 != nil || err2 != nil || err3 != nil {
		return &dbStore{}, errorInitializing
	}
	return &dbStore{db, getByIdStm, getByOriginalUrlStm, upsertUrlStm}, nil
}

func (store *dbStore) getById(id int) (url, error) {
	var u url
	err := store.getByIdStm.QueryRow(id).Scan(&u.id, &u.originalUrl, &u.shortUrl)
	if err != nil {
		return u, errorFailToFind
	}
	return u, nil
}

func (store *dbStore) getByOriginalUrl(originalUrl string) (url, error) {
	var u url
	err := store.getByOriginalUrlStm.QueryRow(originalUrl).Scan(&u.id, &u.originalUrl, &u.shortUrl)
	if err != nil {
		return u, errorFailToFind
	}
	return u, nil
}

func (store *dbStore) upsertUrl(u url) error {
	_, err := store.upsertUrlStm.Query(u.id, u.originalUrl, u.shortUrl, u.shortUrl)
	if err != nil {
		return errorFailToSave
	}
	return nil
}
