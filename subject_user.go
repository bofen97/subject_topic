package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// user subject

// user http post id ， subject  ,customlabelTopic to db

type UserSubjectServer struct {
	Subjt   *SubjectTable
	Session *SessionTable
}
type UserSubjectServerData struct {
	Session     string `json:"session"`
	Topic       string `json:"topic"`
	CustomLabel string `json:"customlabel"`
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

			//then put info to db
			if UsubjData.Topic != "" {
				err := usubj.Subjt.InsertIdTopic(uid, UsubjData.Topic)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				fmt.Printf("user [%d]  subject topic [%s]\n", uid, UsubjData.Topic)
				w.WriteHeader(http.StatusOK)
				return
			} else {
				err := usubj.Subjt.InsertIdCustomLabel(uid, UsubjData.CustomLabel)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				fmt.Printf("user [%d]  subject customlabel [%s]\n", uid, UsubjData.CustomLabel)
				go ToCache(UsubjData.CustomLabel)
				w.WriteHeader(http.StatusOK)
				return

			}

		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusBadRequest)

}

type CacheData struct {
	Topic string `json:"topic"`
}

func ToCache(topic string) error {
	var cachedata = CacheData{
		Topic: topic,
	}

	data, err := json.Marshal(cachedata)
	if err != nil {
		log.Print(err)
		return err
	}
	req, err := http.NewRequest("POST", cacheServer+"/cache", bytes.NewBuffer(data))
	if err != nil {
		log.Print(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return err
	}
	if response.StatusCode == 200 {
		log.Printf("[%s] To Cache \n", topic)
	}
	return nil

}
