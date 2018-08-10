package main

import (
	"net/http"
	"fmt"
	_ "github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"encoding/json"
	"github.com/gorilla/mux"
)

var db *gorm.DB
var err error

func InitialMigration()  {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil{
		fmt.Println(err.Error())
		panic("Fault occured while migrating")
	}
	defer db.Close()
	db.AutoMigrate(&MemberDb{})

}


func AllMembers(w http.ResponseWriter, r *http.Request)  {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil{
		fmt.Println(err.Error())
		panic("Fault occured while migrating")
	}
	defer db.Close()
	var members []MemberDb
	db.Find(&members)
	json.NewEncoder(w).Encode(members)

}

func NewMember(w http.ResponseWriter, r *http.Request){
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil{
		fmt.Println(err.Error())
		panic("Fault occured while migrating")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	phNumber := vars["phNumber"]
	carNumber := vars["carNumber"]

	db.Create(&MemberDb{Name: name,PhoneNumber:phNumber, CarNumber:carNumber})

	fmt.Fprintf(w, "MemberDb created")

}

func DeleteMember(w http.ResponseWriter, r *http.Request)  {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil{
		fmt.Println(err.Error())
		panic("Fault occured while migrating")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var member MemberDb

	db.Where("name = ?", name).Find(&member)
	db.Delete(member)

	fmt.Fprintf(w, "MemberDb deleted")
}