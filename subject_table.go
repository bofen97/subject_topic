package main

import (
	"database/sql"
	"log"
)

type SubjectTable struct {
	db *sql.DB
}

func (sub *SubjectTable) Connect(url string) (err error) {

	sub.db, err = sql.Open("mysql", url)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = sub.db.Ping()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (sub *SubjectTable) CreateTable() error {

	query := `
	CREATE TABLE IF NOT EXISTS userSubjectTable (
		uid INT(11) NOT NULL,
		topic TEXT ,
		customlabel TEXT 
	);`

	if _, err := sub.db.Exec(query); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
func (sub *SubjectTable) InsertIdTopic(uid int, topic string) error {
	if topic == "" {
		return nil
	}
	queryStr := `
	insert into userSubjectTable(uid,topic) values(?,?)
	`

	if _, err := sub.db.Exec(queryStr, uid, topic); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
func (sub *SubjectTable) InsertIdCustomLabel(uid int, custom string) error {
	if custom == "" {
		return nil
	}
	queryStr := `
	insert into userSubjectTable(uid,customlabel) values(?,?)
	`

	if _, err := sub.db.Exec(queryStr, uid, custom); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
