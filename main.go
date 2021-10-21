package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name     string
	Email    string
	Password string
	Mobile   string
	Role     int
	Modules  []string
	UserId   string
}

var users []User


func main() {
	fmt.Println("API🚀")
	r := mux.NewRouter()
	r.HandleFunc("/", standardRoute)
	r.HandleFunc("/user", userRoute_Create).Methods("POST")
	listen := http.ListenAndServe(":5000", r)
	log.Fatal(listen)
}

func standardRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<h1>🐇 golang</h1>`))

}

func userRoute_Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Enter the course details")
	}
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	fmt.Println(user.isEmty())

	if user.isBodyContains() {
		json.NewEncoder(w).Encode("No data present ")
		return
	}
	rand.Seed(time.Now().UnixNano())
	user.UserId = strconv.Itoa(rand.Intn(100))

	hash, _ := passwordHasher(user.Password)

	user.Password = hash

	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

// func handleError(err error) {
// 	if err == nil {
// 		panic(err)
// 	}
// }

func (user *User) isBodyContains() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return user.Name == "" ||user.Email=="" ||user.Password==""
}

func passwordHasher(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	return string(hash), err
}
