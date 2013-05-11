package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Registration struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ErrorResponse struct {
	Status  int8   `json:"status"`
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	ref := r.Referer()

	if ref == "" {
		ref = "localhost"
	}

	fmt.Printf("referrer: \"%s\"", ref)

	db := Db{ref, nil}

	err := db.sCheck()

	if err != nil {
		serveError(w, err.Error())
		return
	}

	switch r.Method {
	case "GET":
		sID := r.URL.Path[1:]
		var id int64 = 0

		if sID == "" {
			id = 0
		} else {
			convid, err := strconv.ParseInt(sID, 10, 64)
			id = convid

			if err != nil {
				serveError(w, err.Error())
				return
			}
		}

		if id == 0 {
			err = db.List()
			if err != nil {
				serveError(w, err.Error())
				return
			}
		} else {
			r, err := db.Get(id)

			if err != nil {
				serveError(w, err.Error())
				return
			}

			enc := json.NewEncoder(w)
			w.Header().Set("Content-Type", "application/json")
			enc.Encode(r)
		}

	case "POST":
		dec := json.NewDecoder(r.Body)

		r := new(Registration)

		err := dec.Decode(&r)

		if err != nil {
			serveError(w, err.Error())
			return
		}

		r.Password = Encrypt(r.Password)

		r, err = db.Create(r)

		if err != nil {
			serveError(w, err.Error())
			return
		}

		enc := json.NewEncoder(w)
		w.Header().Set("Content-Type", "application/json")
		enc.Encode(r)
	case "PUT":
		dec := json.NewDecoder(r.Body)

		r := new(Registration)

		err := dec.Decode(&r)

		if err != nil {
			serveError(w, err.Error())
			return
		}

		r.Password = Encrypt(r.Password)

		r, err = db.Update(r)

		if err != nil {
			serveError(w, err.Error())
			return
		}

		enc := json.NewEncoder(w)
		w.Header().Set("Content-Type", "application/json")
		enc.Encode(r)
	case "DELETE":
		sID := r.URL.Path[1:]
		id, err := strconv.ParseInt(sID, 10, 64)

		err = db.Delete(id)

		if err != nil {
			serveError(w, err.Error())
			return
		}
		serveOK(w, "registration deleted")

	}
}

func serveError(w http.ResponseWriter, msg string) {
	//var err = errors.New(msg)
	w.WriteHeader(http.StatusInternalServerError)
	//w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	status := ErrorResponse{
		Status:  -1,
		Message: msg,
	}

	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(status)
	return
}
func serveOK(w http.ResponseWriter, msg string) {
	status := ErrorResponse{
		Status:  1,
		Message: msg,
	}

	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(status)
}
