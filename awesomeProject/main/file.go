package main

import (
					"log"
			"net/http"
	"github.com/gorilla/mux"
	"os"
			"encoding/csv"
	"fmt"
	)

func CreateMemberInFile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)
	name := vars["name"]
	phNumber := vars["phNumber"]
	carNumber := vars["carNumber"]

	member := []string{name, phNumber, carNumber}
	var fileWriteErr string

	file, err := os.OpenFile("Members.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if file != nil {
		checkError("Cannot create file", err)
		defer file.Close()
		fileWriteErr = writeFile(file, member)
	} else {
		os.Create("Members.csv")
		checkError("Cannot create file", err)
		file.Close()
		file, err := os.OpenFile("Members.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		defer file.Close()
		checkError("Cannot create file", err)
		fileWriteErr = writeFile(file, member)
	}
	if fileWriteErr != "OK" {
		fmt.Fprintf(w, "Error occured whole writing file")
	}
}

func checkError(message string, err error)  {
	if err != nil{
		log.Fatal(message, err)
	}

}

func writeFile(file *os.File, member []string) (err string) {
	writer := csv.NewWriter(file)
	defer writer.Flush()

	errW := writer.Write(member)

	checkError("Cannot write file", errW)

	if errW != nil{
		log.Fatal("Cannot write file", err)
		return "Cannot write file"
	} else {
		return "OK"
	}

}

func ReadAllMembersInFile(w http.ResponseWriter, r *http.Request)  {

	file, err := os.Open("Members.csv")
	defer file.Close()
	checkError("Cannot create file", err)

	reader := csv.NewReader(file)
	members, err := reader.ReadAll()

	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, members)
}

func FindMemberInFileByPhoneNumber(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)
	phNumber := vars["phNumber"]

	file, err := os.Open("Members.csv")
	defer file.Close()
	checkError("Cannot create file", err)

	reader := csv.NewReader(file)
	members, err := reader.ReadAll()

	if err != nil {
		fmt.Fprintln(w, err)
	}

	var memberExc string = findMemberInFileByColumnVAlue(members, phNumber, 1)

	if memberExc !="OK" {
		fmt.Fprintln(w, memberExc)
	} else {
		fmt.Fprintln(w, memberExc)
	}


}

func findMemberInFileByColumnVAlue(records [][]string, val string, col int) string {
	for i, row := range records {
		if row[col] == val {
			return "OK"
		}
		i++
	}
	return "No such person"
}