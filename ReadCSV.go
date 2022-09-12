package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type Course struct {
	assignedFlag       string
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
	Assignment_Date    string
	dueDate            string
	completeDate       string
	ulearn_domain      string
}

var ALLCourses []Course
var csvTomap map[string][]Course

func main() {
	f, err := os.Open("Dataset.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	//data, err := csvReader.ReadAll()
	_, err1 := csvReader.Read()
	if err1 != nil {
		log.Fatal(err1)
	}
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(data)

	csvTomap = make(map[string][]Course)
	for {
		line, err := csvReader.Read()

		//fmt.Print("Line Contains", line)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("error reading file", err)
		}
		_, isPresent := csvTomap[line[3]]
		if isPresent {
			csvTomap[line[3]] = append(csvTomap[line[3]], Course{
				assignedFlag:       line[0],
				TableNames:         line[1],
				ulearn_countryName: line[2],
				Win_Number:         line[3],
				subCourse:          line[4],
				activityId:         line[5],
				userEmail:          line[6],
				fullName:           line[7],
				userid:             line[8],
				enrollType:         line[9],
				jobType:            line[10],
				ManagerEmail:       line[11],
				ManagerFullName:    line[12],

				Manager_userid:  line[13],
				activityName:    line[14],
				categoryType:    line[15],
				Assignment_Date: line[16],
				dueDate:         line[17],

				completeDate:  line[18],
				ulearn_domain: line[18],
			})

		} else {

			csvTomap[line[3]] = append(csvTomap[line[3]], Course{
				assignedFlag:       line[0],
				TableNames:         line[1],
				ulearn_countryName: line[2],
				Win_Number:         line[3],
				subCourse:          line[4],
				activityId:         line[5],
				userEmail:          line[6],
				fullName:           line[7],
				userid:             line[8],
				enrollType:         line[9],
				jobType:            line[10],
				ManagerEmail:       line[11],
				ManagerFullName:    line[12],
				Manager_userid:     line[13],
				activityName:       line[14],
				categoryType:       line[15],
				Assignment_Date:    line[16],
				dueDate:            line[17],
				completeDate:       line[18],
				ulearn_domain:      line[18],
			})
		}
	}

	fmt.Println(csvTomap)
}
