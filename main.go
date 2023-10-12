package main

import (
	"log"
	"net/http"
	"os"

	serviceRegisterCenter "github.com/bofen97/ServiceRegister"
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
	etcdserver := os.Getenv("etcdserver")
	if etcdserver == "" {
		log.Fatal("etcdserver is none")
		return
	}
	src, err := serviceRegisterCenter.NewRegisteService([]string{
		etcdserver,
	}, 5)
	if err != nil {
		log.Fatal(err)
	}
	go src.PutServiceAddr("subject_topic", "subject_topic"+serverPort)
	go src.ListenLaser()

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
	mux.HandleFunc("/cancel", usubj.CancelSubjectTopic)
	server := &http.Server{
		Addr:    serverPort,
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
