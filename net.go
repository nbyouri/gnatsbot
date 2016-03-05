package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
)

var bug struct {
	Number      string
	Category    string
	Synopsis    string
	Severity    string
	Priority    string
	Responsible string
	Class       string
	Originator  string
}

/* Print new bugs */
func main() {
	//var url string = "http://mail-index.netbsd.org/netbsd-bugs/2016/03/01/msg045099.html"
	//var body = getDataFromURL(url)
	file, err := os.Open("bug.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println("RUNNING:")

	parseBug(file)

	fmt.Println(bug.Number, bug.Category, bug.Synopsis, bug.Severity,
		bug.Priority, bug.Responsible, bug.Class, bug.Originator)

	return
}

//func getDataFromURL(url string) (body []byte) {
//	resp, err := http.Get(url)
//	if err != nil {
//		fmt.Println("Failed to get ", url)
//		return
//	}
//	defer resp.Body.Close()
//	body, err = ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	return
//}

// Read bug file line by line and extract needed information
func parseBug(f *os.File) {
	s := bufio.NewReader(f)

	ref := reflect.ValueOf(&bug).Elem()

	numFieldsParsed := 0
	for {
		if numFieldsParsed == ref.NumField() {
			return
		}
		line, err := s.ReadString(10)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Panic("Failed to read bug")
		}

		if strings.HasPrefix(line, ">") {
			for i := 0; i < ref.NumField(); i++ {
				fieldName := ref.Type().Field(i).Name
				if strings.Contains(line, fieldName) {
					value := strings.Join(strings.Fields(line)[1:], " ")
					if len(value) > 100 {
						log.Panic("Malformed bug file")
					}
					ref.FieldByName(fieldName).SetString(value)
					numFieldsParsed++
				}
			}
		}
	}
}
