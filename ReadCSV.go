package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	m := CSVFileToMap()
	fmt.Println(m)
}

func CSVFileToMap() (returnMap map[string]map[string]string) {
	var data = map[string]map[string]string{}

	filePath := "data.csv"
	// read csv file
	csvfile, err := os.Open(filePath)
	if err != nil {
		return nil
	}

	defer csvfile.Close()
	csvfile.Seek(0, 0)
	reader := csv.NewReader(csvfile)

	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		return nil
	}
	x := []string{}
	for i := 1; i < len(rawCSVdata); i++ {
		x = append(x, rawCSVdata[i][23])

	}
	fmt.Println("Raw CSVVVV", len(rawCSVdata[0]))
	//fmt.Printf("Type %T", rawCSVdata)
	header := []string{} // holds first row (header)
	//x := []string{"key1", "key2", "key3"}
	//fmt.Println("xxxxxx", x)
	//fmt.Println("Record", record)
	//win := "col1"
	//fmt.Println("Header", header)

	for lineNum, record := range rawCSVdata {
		//fmt.Println("Line number", lineNum)
		// for first row, build the header slice
		if lineNum == 0 {
			for i := 0; i < len(rawCSVdata[0]); i++ {
				//fmt.Println("Header", header)
				header = append(header, strings.TrimSpace(record[i]))

			}
			//fmt.Println("Headder", header)
		} else {

			data[x[lineNum-1]] = map[string]string{}

			for j := 0; j < len(header); j++ {

				data[x[lineNum-1]][header[j]] = rawCSVdata[lineNum][j]

			}

			// returnMap = append(returnMap, data)
			//fmt.Println("\n Return Map", data)

		}

	}
	//fmt.Println("Dataaaaaaaaa", data)
	count := 0
	var data1 = ""
	for key, ele := range data {
		//fmt.Println("Key:", key)
		if key == "" {
			fmt.Println("Element", ele)
			data1 = fmt.Sprint(ele)
			count++
		}
	}
	f, err := os.Create("outputt.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	//data1 := fmt.Sprint(data)
	_, err2 := f.WriteString(fmt.Sprint(data1))

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
	fmt.Println("Number of Empty Records :: ", count)
	return data

}
