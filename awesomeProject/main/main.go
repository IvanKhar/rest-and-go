package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"log"
)




func helloWorld(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello World")
}

func handleRequest()  {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", helloWorld)
	router.HandleFunc("/members", AllMembers).Methods("GET")
	router.HandleFunc("/membersFile", ReadAllMembersInFile).Methods("GET")
	router.HandleFunc("/memberFile/{phNumber}", FindMemberInFileByPhoneNumber).Methods("GET")
	router.HandleFunc("/member/{name}/{phNumber}/{carNumber}", NewMember).Methods("POST")
	router.HandleFunc("/members/{name}/{phNumber}/{carNumber}", CreateMemberInFile).Methods("POST")
	router.HandleFunc("/memberDel/{name}", DeleteMember).Methods("DELETE")
	//router.HandleFunc("/members", AllMembers).Methods("GET")
	//router.HandleFunc("/member", allMembers).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main()  {
	fmt.Println("Hello")

	InitialMigration()

	handleRequest()
}