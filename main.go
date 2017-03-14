package main

import (
	"crypto/md5"

	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
)

// cookie handling
var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func getUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

func setSession(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

// Login handler

func loginHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("email")
	pass := GetMD5Hash(request.FormValue("password"))
	db, err := sql.Open("mysql", "root:@/cyza?charset=utf8")
	rows, err := db.Query("SELECT id,username,email FROM users WHERE email='" + name + "'  and password='" + pass + "'")
	checkErr(err)
	user := []string{}
	for rows.Next() {
		var id int
		var username string
		var email string

		err = rows.Scan(&id, &username, &email)
		checkErr(err)
		user = append(user, username)

	}

	redirectTarget := "/"
	if len(user) != 0 {
		// .. check credentials ..
		setSession(name, response)
		redirectTarget = "/profil"
	}
	http.Redirect(response, request, redirectTarget, 302)
}

// logout handler

func logoutHandler(response http.ResponseWriter, request *http.Request) {
	clearSession(response)
	http.Redirect(response, request, "/", 302)
}

func indexPageHandler(response http.ResponseWriter, request *http.Request) {
	//fmt.Fprintf(response, indexPage)
	t, _ := template.ParseFiles("home.gtpl")
	t.Execute(response, nil)
}

func internalPageHandler(response http.ResponseWriter, request *http.Request) {
	userName := getUserName(request)
	if userName != "" {
		db, err := sql.Open("mysql", "root:@/cyza?charset=utf8")
		rows, err := db.Query("SELECT email,full_name,phone,address,created FROM users WHERE email='" + userName + "'")
		checkErr(err)
		user := []string{}
		for rows.Next() {
			var email string
			var full_name string
			var phone string
			var address string
			var created string

			err = rows.Scan(&email, &full_name, &phone, &address, &created)
			checkErr(err)
			user = append(user, email)
			user = append(user, full_name)
			user = append(user, phone)
			user = append(user, address)
			user = append(user, created)

		}
		t, _ := template.ParseFiles("dash.gtpl")
		t.Execute(response, user)
		//fmt.Fprintf(response, internalPage, userName)
	} else {
		http.Redirect(response, request, "/", 302)
	}
}

// Register handler
func registerHandler(response http.ResponseWriter, request *http.Request) {
	var msg string
	if request.Method == "GET" {

	} else {

		name := request.FormValue("name")
		password := GetMD5Hash(request.FormValue("password"))
		email := request.FormValue("email")
		db, err := sql.Open("mysql", "root:@/cyza?charset=utf8")

		rows, err := db.Query("SELECT email FROM users WHERE  email='" + email + "'")
		checkErr(err)
		user := []string{}
		for rows.Next() {
			var _email string
			err = rows.Scan(&_email)
			checkErr(err)
			user = append(user, _email)

		}
		if len(user) != 0 {
			msg = "User exists"

		} else {
			stmt, err := db.Prepare("INSERT INTO users SET username=? , password=? , email=? ")
			checkErr(err)
			stmt.Exec(name, password, email)

			msg = "User added"
		}
	}
	t, _ := template.ParseFiles("register.gtpl")
	t.Execute(response, msg)
}

//Lost Password Handler
func lostHandler(response http.ResponseWriter, request *http.Request) {
	var msg string

	if request.Method == "GET" {

	} else {
		email := request.FormValue("email")
		db, err := sql.Open("mysql", "root:@/cyza?charset=utf8")
		rows, err := db.Query("SELECT email,password FROM users WHERE  email='" + email + "'")
		checkErr(err)
		user := []string{}
		for rows.Next() {
			var _email string
			var _password string
			err = rows.Scan(&_email, &_password)
			checkErr(err)
			user = append(user, _email)
			user = append(user, _password)

		}
		if len(user) != 0 {
			h := GetMD5Hash(email)
			msg = "Email Sent to : " + email
			path := "http://localhost:8000/reset?token=" + h
			fmt.Println(path)
			db.Query("UPDATE users SET reset_hash='" + h + "'  WHERE   email='" + email + "'")

			send(path, email)
		} else {
			msg = "Email don't exists : "
		}

	}
	t, _ := template.ParseFiles("lost.gtpl")
	t.Execute(response, msg)
}
func resetHandler(response http.ResponseWriter, request *http.Request) {
	msg := ""
	token := request.FormValue("token")
	pass := request.FormValue("password")
	if request.Method == "GET" {

	} else {
		if token != "" && pass != "" {
			db, err := sql.Open("mysql", "root:@/cyza?charset=utf8")
			db.Query("UPDATE users SET password='" + pass + "'  WHERE   reset_hash='" + token + "'")
			checkErr(err)
			http.Redirect(response, request, "/", 302)
		}

	}
	t, _ := template.ParseFiles("forgot.gtpl")
	t.Execute(response, msg)
}
func profilHandler(response http.ResponseWriter, request *http.Request) {
	userName := getUserName(request)
	if userName != "" {
		if request.Method == "GET" {

		} else {
			full_name := request.FormValue("full_name")
			address := request.FormValue("address")
			phone := request.FormValue("phone")

			db, err := sql.Open("mysql", "root:@/cyza?charset=utf8")
			checkErr(err)
			db.Query("UPDATE users SET full_name='" + full_name + "' , address='" + address + "' , phone='" + phone + "' WHERE   email='" + userName + "'")
			//checkErr(err)
			http.Redirect(response, request, "/internal", 302)
		}
		t, _ := template.ParseFiles("profile.gtpl")
		t.Execute(response, nil)

	} else {
		http.Redirect(response, request, "/", 302)
	}
}

// server main method

var router = mux.NewRouter()

func main() {

	router.HandleFunc("/", indexPageHandler)
	router.HandleFunc("/internal", internalPageHandler)
	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/register", registerHandler).Methods("POST", "GET")
	router.HandleFunc("/lost", lostHandler).Methods("POST", "GET")
	router.HandleFunc("/reset", resetHandler).Methods("POST", "GET")
	router.HandleFunc("/profil", profilHandler).Methods("POST", "GET")
	router.HandleFunc("/logout", logoutHandler).Methods("POST")

	router.HandleFunc("/api/login", ApiLoginHandler).Methods("POST", "GET")
	router.HandleFunc("/api/register", ApiRegisterHandler).Methods("POST", "GET")

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/", router)
	http.ListenAndServe(":8000", nil)
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func send(body string, to string) {
	from := "chiheb.design@gmail.com"
	pass := ""

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Your Password\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:25",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

}
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

type Message struct {
	Name  string `json:"name"`
	Error bool   `json:"error"`
}

// API Login

func ApiLoginHandler(response http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	pass := request.FormValue("password")
	fmt.Println(email)
	fmt.Println(pass)
	db, err := sql.Open("mysql", "root:@/cyza?charset=utf8")
	rows, err := db.Query("SELECT id,username,email FROM users WHERE email='" + email + "'  and password='" + pass + "'")
	checkErr(err)
	user := []string{}
	for rows.Next() {
		var id int
		var username string
		var email string

		err = rows.Scan(&id, &username, &email)
		checkErr(err)
		user = append(user, username)
		fmt.Println(id)

	}
	if len(user) != 0 {
		if err := json.NewEncoder(response).Encode(Message{Name: "Logged", Error: false}); err != nil {
			panic(err)
		}

	} else {
		if err := json.NewEncoder(response).Encode(Message{Name: "Error : Bad request", Error: true}); err != nil {
			panic(err)
		}
	}

}

// API Register
func ApiRegisterHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	password := request.FormValue("password")
	email := request.FormValue("email")

	if name != "" && password != "" && email != "" {
		db, err := sql.Open("mysql", "root:@/cyza?charset=utf8")
		rows, err := db.Query("SELECT email FROM users WHERE  email='" + email + "'")
		checkErr(err)
		user := []string{}
		for rows.Next() {
			var _email string
			err = rows.Scan(&_email)
			checkErr(err)
			user = append(user, _email)

		}
		if len(user) != 0 {
			if err := json.NewEncoder(response).Encode(Message{Name: "User Exists", Error: true}); err != nil {
				panic(err)
			}

		} else {

			stmt, err := db.Prepare("INSERT INTO users SET username=? , password=? , email=? ")
			checkErr(err)
			stmt.Exec(name, password, email)

			if err := json.NewEncoder(response).Encode(Message{Name: "User added", Error: false}); err != nil {
				panic(err)
			}

		}
	} else {
		if err := json.NewEncoder(response).Encode(Message{Name: "Error : Bad request", Error: true}); err != nil {
			panic(err)
		}
	}
}

/*
*
*
*
*
*
*
*
*
*
 */
