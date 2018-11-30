package service

import(
	"encoding/json"
	"log"
	"fmt"
	"strings"
	
	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

const(
	SecretKey = "secret"
)

func fatal(err error){
	if err!=nil{
		log.Fatal(err)
	}
}

type UserCredentials struct{
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}

type Token struct{
	Token string	`json:"token"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request){
	var user UserCredentials
	
	err := json.NewDecoder(r.Body).Decode(&user)
	if err!=nil{
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w,"Error in request")
		return
	}

	result:=DBClient.SelectRecord(SecretKey,strings.ToLower(user.Username)	
	if len(result)==0{
		w.WriteHeader(http.StatusForbidden)
		fmt.Println("Error logging in")
		fmt.Fprintf(w,"Invalid credentials")
		return
	}


}

