package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type ALLRecord struct { //declare struct
	Assignment_Date    string
	completeDate       string
	dueDate            string
	flowrunKey         string
	assignedFlag       string
	ISlateFlag         string
	IncompleteFlag     string
	completeFlag       string
	lateFlag           string
	assignmentFY       string
	assignedQtr        string
	dueFY              string
	dueQtr             string
	completedFY        string
	completedQtr       string
	trainingKey        string
	flowrunMonth       string
	flowrunFY          string
	flowrunQtr         string
	dateTimestamp      string
	completionStatus   string
	TableNames         string
	ulearn_countryName string
	Win_Number         string
	subCourse          string
	activityId         string
	userEmail          string
	fullName           string
	userid             string
	enrollType         string
	jobType            string
	ManagerEmail       string
	ManagerFullName    string
	Manager_userid     string
	activityName       string
	categoryType       string
	ulearn_domain      string
}

func createList(data [][]string) []ALLRecord {
	var List []ALLRecord
	for i, line := range data {
		if i > 0 { // omit header line
			var rec ALLRecord
			for j, field := range line {
				if j == 0 {
					rec.Assignment_Date = field
				} else if j == 1 {
					rec.completeDate = field
				} else if j == 2 {
					rec.dueDate = field
				} else if j == 3 {
					rec.flowrunKey = field
				} else if j == 4 {
					rec.assignedFlag = field
				} else if j == 5 {
					rec.ISlateFlag = field
				} else if j == 6 {
					rec.IncompleteFlag = field
				} else if j == 7 {
					rec.completeFlag = field
				} else if j == 8 {
					rec.lateFlag = field
				} else if j == 9 {
					rec.assignmentFY = field
				} else if j == 10 {
					rec.assignedQtr = field
				} else if j == 11 {
					rec.dueFY = field
				} else if j == 12 {
					rec.dueQtr = field
				} else if j == 13 {
					rec.completedFY = field
				} else if j == 14 {
					rec.completedQtr = field
				} else if j == 15 {
					rec.trainingKey = field
				} else if j == 16 {
					rec.flowrunMonth = field
				} else if j == 17 {
					rec.flowrunFY = field
				} else if j == 18 {
					rec.flowrunQtr = field
				} else if j == 19 {
					rec.dateTimestamp = field
				} else if j == 20 {
					rec.completionStatus = field
				} else if j == 21 {
					rec.TableNames = field
				} else if j == 22 {
					rec.ulearn_countryName = field
				} else if j == 23 {
					rec.Win_Number = field
				} else if j == 24 {
					rec.subCourse = field
				} else if j == 25 {
					rec.activityId = field
				} else if j == 26 {
					rec.userEmail = field
				} else if j == 27 {
					rec.fullName = field
				} else if j == 28 {
					rec.userid = field
				} else if j == 29 {
					rec.enrollType = field
				} else if j == 30 {
					rec.jobType = field
				} else if j == 31 {
					rec.ManagerEmail = field
				} else if j == 32 {
					rec.ManagerFullName = field
				} else if j == 33 {
					rec.Manager_userid = field
				} else if j == 34 {
					rec.activityName = field
				} else if j == 35 {
					rec.categoryType = field
				} else if j == 36 {
					rec.ulearn_domain = field
				}

			}

			List = append(List, rec)
		}

	}
	return List
}

func main() {

	f, err := os.Open("data.csv") //Open CSV file
	if err != nil {
		log.Fatal(err)
	}

	// close the file
	defer f.Close()

	// reading CSV File
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// convert records to array of structs
	// data = data[:500]
	// fmt.Println("Dataaaaaaaaa \n", data)
	// print the array
	List := createList(data)

	// print the array
	fmt.Printf("%+v\n", List)
}
