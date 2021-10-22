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


func (user *User) isEmty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return user.Name == ""
}
func (user *User) isAdmin() bool {
	return user.Role == 0
}

func main() {

	fmt.Println("API🚀")
	r := mux.NewRouter()
	r.HandleFunc("/", standardRoute)
	r.HandleFunc("/user", userRoute_Create).Methods("POST")
	r.HandleFunc("/users", getAllUsers_Admin).Methods("GET")
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


	if user.isEmty() {
		json.NewEncoder(w).Encode("Some of the fileds is Missing")
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

func getAllUsers_Admin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Admin route_GetAllUsers")
	w.Header().Set("Content-Type", "application/json")

	defer r.Body.Close()
	var user User

	fmt.Println("data", user.isAdmin())

	if r.Body != nil {

		fmt.Println("Admin")
		json.NewEncoder(w).Encode(users)

	} else {
		json.NewEncoder(w).Encode("No users so far")
	}

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
