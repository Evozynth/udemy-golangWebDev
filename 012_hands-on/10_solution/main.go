package main

import (
	"text/template"
	"log"
	"encoding/csv"
	"os"
	"time"
	"strconv"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type Record struct{
	Date time.Time
	Close float64
}

func main() {
	http.HandleFunc("/", serve)
	http.ListenAndServe(":8080", nil)
}


func serve(res http.ResponseWriter, req *http.Request) {
	f, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	rows, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	csvData := make([]Record, 0, len(rows))
	for i, row := range rows {
		if i == 0 {
			continue
		}
		date, _ := time.Parse("2006-01-02", row[0])
		closing, _ := strconv.ParseFloat(row[4], 64)

		csvData = append(csvData, Record{
			Date: date,
			Close: closing,
		})
	}

	tpl.Execute(res, csvData)
}
