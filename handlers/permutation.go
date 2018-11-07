package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/satori/go.uuid"

	"permutation-server/models"
	"permutation-server/pkg/database"
)

const maxBodySize = 1048576

func Init(w http.ResponseWriter, r *http.Request) {
	bodyStream := http.MaxBytesReader(w, r.Body, maxBodySize)
	body, err := ioutil.ReadAll(bodyStream)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"invalid body"}`))
		return
	}
	arr := make([]int64, 0)
	err = json.Unmarshal(body, &arr)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"invalid body"}`))
		return
	}
	uid, err := uuid.NewV4()
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"internal error"}`))
		return
	}
	cookie := &http.Cookie{
		Name:  "uid",
		Value: uid.String(),
		Path:  "/",
	}
	http.SetCookie(w, cookie)

	perm := models.Permutation{
		Uuid: uid.String(),
		Data: arr,
	}
	db, err := database.NewDB()
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"internal error"}`))
		return
	}
	if err := db.Save(&perm); err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"internal error"}`))
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}

func Next(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("uid")
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message":"auth error"}`))
		return
	}
	uuid := cookie.Value
	db, err := database.NewDB()
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"internal error"}`))
		return
	}
	perm, err := db.Find(uuid)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message":"array does not init"}`))
		return
	}
	rvsIdx := perm.NextPermutation()
	if rvsIdx != -1 {
		if err := db.Update(*perm, rvsIdx); err != nil {
			log.Print(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"internal error"}`))
			return
		}
	}
	res, err := json.Marshal(perm.Data)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "internal error"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
