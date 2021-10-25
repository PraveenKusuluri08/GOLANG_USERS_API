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
	r.HandleFunc("/user/{id}", getSingleUser).Methods("GET")
	r.HandleFunc("/user/{id}", UpdateSingleUser).Methods("PUT")
	r.HandleFunc("/user/{id}", deleteSingleUser).Methods("DELETE")
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
		json.NewEncoder(w).Encode("Some of the fields is missing")
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

	if r.Body != nil && user.isAdmin() {

		fmt.Println("Admin")
		json.NewEncoder(w).Encode(users)

	} else {
		json.NewEncoder(w).Encode("No users so far")
	}

}

func getSingleUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get single User document")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params["id"])

	for _, user := range users {
		if user.UserId == params["id"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode("No user found with given id")
}

// by this end point data is overrided with
//because we are just omitting data from slice and
//making changes in the data and adding data again
//into the slice if user changes only one field
//data is totally overrided with new data
// means with empty values we control this flow in frontend =>react
func UpdateSingleUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Single User")
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()
	params := mux.Vars(r)

	for id, user := range users {
		if user.UserId == params["id"] {
			users = append(users[:id], users[id+1:]...)
			var user User
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.UserId = params["id"]
			users = append(users, user)
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	json.NewEncoder(w).Encode("User is not there with the given id")

}

func deleteSingleUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete single user")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for id, user := range users {
		if user.UserId == params["id"] {
			users = append(users[:id], users[id+1:]...)
			json.NewEncoder(w).Encode(&user)
			break
		}
	}
	json.NewEncoder(w).Encode("massage:Course is deleted successfully")
}

// func handleError(err error) {
// 	if err == nil {
// 		panic(err)
// 	}
// }

func passwordHasher(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	return string(hash), err
}

//helper to check user exist or not

func UserExistsOrNot(userId string) bool {
	var user User
	return user.UserId == userId
}
