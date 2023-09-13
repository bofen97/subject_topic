package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	sqlurl := os.Getenv("sqlurl")
	if sqlurl == "" {
		log.Fatal("sqlurl is none")
		return
	}
	serverPort := os.Getenv("serverport")
	if serverPort == "" {
		log.Fatal("serverPort is none")
		return
	}

	usubj := new(UserSubjectServer)
	usubj.Subjt = new(SubjectTable)
	usubj.Session = new(SessionTable)

	if err := usubj.Subjt.Connect(sqlurl); err != nil {
		log.Fatal(err)
	}
	if err := usubj.Subjt.CreateTable(); err != nil {
		log.Fatal(err)
	}
	log.Print("Created table !")

	if err := usubj.Session.Connect(sqlurl); err != nil {
		log.Fatal(err)
	}
	log.Print("Session connect !")

	mux := http.NewServeMux()
	mux.Handle("/subject", usubj)
	server := &http.Server{
		Addr:    serverPort,
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
