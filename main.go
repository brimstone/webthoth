package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var rooms map[string]Room

func helloFunc(res http.ResponseWriter, req *http.Request) {
	log.Println("handling /")
	fmt.Fprintln(res, "hello!")
}

func roomFunc(res http.ResponseWriter, req *http.Request) {
	r := strings.Split(req.RequestURI, "/")
	// Check to make sure we have enough arguments
	roomid := r[2]
	if roomid == "" {
		return
	}
	id := ""
	if len(r) == 4 {
		id = r[3]
	}
	log.Println("handling room", roomid, id)
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, X-Registry-Auth")
	res.Header().Add("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
	// new means we're looking for an offer
	if id == "" {
		// List Description IDs for this room, if any
		if req.Method == "GET" {
			if _, ok := rooms[roomid]; !ok {
				res.Header().Set("Content-Type", "application/json")
				fmt.Fprint(res, "[]")
				return
			}
			output, _ := json.Marshal(rooms[roomid].ListDescriptions())
			res.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(res, string(output))
			// Add a description ID to this room, return the ID
		} else if req.Method == "POST" {
			if _, ok := rooms[roomid]; !ok {
				rooms[roomid] = NewRoom()
			}
			payload, err := ioutil.ReadAll(req.Body)
			if err != nil {
				fmt.Fprint(res, "Got error", err)
				return
			}
			key, err := rooms[roomid].NewDescription(string(payload))
			fmt.Fprint(res, key)
		}
	} else {
		// we're trying to get a specific ID
		if req.Method == "GET" {
			if _, ok := rooms[roomid]; !ok {
				log.Println("room", roomid, "doesn't exist")
				http.Error(res, "room doesn't exist", 404)
				return
			}
			description, err := rooms[roomid].GetDescription(id)
			if err != nil {
				http.Error(res, err.Error(), 404)
			}
			fmt.Fprint(res, description)
			// we're trying to answer a specifid ID
		} else if req.Method == "POST" {
			if _, ok := rooms[roomid]; !ok {
				log.Println("room", roomid, "doesn't exist")
				http.Error(res, "room doesn't exist", 500)
				return
			}
			payload, err := ioutil.ReadAll(req.Body)
			if err != nil {
				fmt.Fprint(res, "Got error", err)
				return
			}
			err = rooms[roomid].AnswerDescription(id, string(payload))
			if err != nil {
				http.Error(res, err.Error(), 500)
				return
			}
			fmt.Fprint(res, "")
		}
	}
}

func init() {
	rooms = make(map[string]Room)
}

func main() {
	http.HandleFunc("/room/", roomFunc)
	http.HandleFunc("/", helloFunc)
	log.Println("Starting server on port", os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}

}
