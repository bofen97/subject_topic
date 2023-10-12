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
		topic TEXT 
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

func (sub *SubjectTable) DeleteIdTopic(uid int, topic string) error {
	if topic == "" {
		return nil
	}
	queryStr := `
	delete from userSubjectTable where uid=? and topic=?
	`

	if _, err := sub.db.Exec(queryStr, uid, topic); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
