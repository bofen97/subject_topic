package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// user subject

// user http post id ， subject  ,customlabelTopic to db

type UserSubjectServer struct {
	Subjt          *SubjectTable
	Session        *SessionTable
	SubjectService *SubjectServiceTable
}
type UserSubjectServerData struct {
	Session string `json:"session"`
	Topic   string `json:"topic"`
}

func (usubj *UserSubjectServer) CancelSubjectTopic(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		if r.Header.Get("Content-Type") == "application/json" {
			data, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			var UsubjData UserSubjectServerData

			err = json.Unmarshal(data, &UsubjData)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			//check session id alive time .
			uid, err := usubj.Session.QuerySessionAndRetUid(UsubjData.Session)
			if err != nil {
				log.Printf("Session [%s] error  \n", UsubjData.Session)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			//then delete topic code .

			err = usubj.Subjt.DeleteIdTopic(uid, UsubjData.Topic)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			fmt.Printf("user [%d]  cancel topic [%s]\n", uid, UsubjData.Topic)
			w.WriteHeader(http.StatusOK)
			return

		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (usubj *UserSubjectServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		if r.Header.Get("Content-Type") == "application/json" {
			data, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			var UsubjData UserSubjectServerData

			err = json.Unmarshal(data, &UsubjData)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			//check session id alive time .
			uid, err := usubj.Session.QuerySessionAndRetUid(UsubjData.Session)
			if err != nil {
				log.Printf("Session [%s] error  \n", UsubjData.Session)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			yes, err := usubj.SubjectService.UidIsExpires(uid)
			if err != nil {
				log.Print("UidIsExpires  error ")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if yes {
				//reject
				log.Printf("Uid %d has Is Expired .", uid)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			//then put info to db

			err = usubj.Subjt.InsertIdTopic(uid, UsubjData.Topic)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			fmt.Printf("user [%d]  subject topic [%s]\n", uid, UsubjData.Topic)
			w.WriteHeader(http.StatusOK)
			return

		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusBadRequest)

}
