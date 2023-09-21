// 'print(err)' has been added to catch errors
package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"io/fs"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB // variable for database
var directory string = "C:\\Users\\tanma\\Downloads\\NewYorkYellowTaxiData"

func ReadFile(wg *sync.WaitGroup, chan1 chan fs.FileInfo, db *sql.DB) {
	defer wg.Done()

	for {
		files, ok := <-chan1
		if !ok {
			break
		}
		// infinite loop only ends when there are no more files to be opened
		file, err := os.Open(directory + "\\" + files.Name())
		if err != nil {
			fmt.Println("error opening file", err)
			continue
		}

		// csv reader to read a file in csv format
		reader := csv.NewReader(file)
		reader.Comma = ','
		rec, err3 := reader.Read()
		if err3 != nil {
			fmt.Println("Could not read records", err3)
			continue
		}

		// SQL query (without values for now)
		query := "INSERT INTO Yellow_Tripdata (vendorid, tpep_pickup_datetime, tpep_dropoff_datetime, passenger_count, trip_distance, ratecodeid, store_and_fwd_flag, pulocationid, dolocationid, payment_type, fare_amount, extra, mta_tax, tip_amount, tolls_amount, improvement_surcharge, total_amount) VALUES "
		tempQuery := ""
		counter := 0

		for err3 != io.EOF {
			rec, err = reader.Read()
			if err == io.EOF {
				break
			}
			// avoid reading header
			if string(rec[0]) == "VendorID" {
				continue
			}
			if err != nil {
				fmt.Println("Could not read the taxi data file", err)
				continue
			}
			if tempQuery != "" {
				// if tempquery is not empty, insert the values from the rec variable which reads line by line
				queryVal := "(" + rec[0] + "," + "'" + rec[1] + "'" + "," + "'" + rec[2] + "'" + "," + rec[3] + "," + rec[4] + "," + rec[5] + "," + "'" + rec[6] + "'" + "," + rec[7] + "," + rec[8] + "," + rec[9] + "," + rec[10] + "," + rec[11] + "," + rec[12] + "," + rec[13] + "," + rec[14] + "," + rec[15] + "," + rec[16] + ")"
				tempQuery = tempQuery + "," + queryVal
			} else {
				queryVal := "(" + rec[0] + "," + "'" + rec[1] + "'" + "," + "'" + rec[2] + "'" + "," + rec[3] + "," + rec[4] + "," + rec[5] + "," + "'" + rec[6] + "'" + "," + rec[7] + "," + rec[8] + "," + rec[9] + "," + rec[10] + "," + rec[11] + "," + rec[12] + "," + rec[13] + "," + rec[14] + "," + rec[15] + "," + rec[16] + ")"
				tempQuery += queryVal
			}

			if counter%800 == 0 {
				// putting the query and the values together
				// inserting 800 records at once
				tempQuery = query + tempQuery
				//executing the query
				_, err2 := db.Exec(tempQuery)
				if err2 != nil {
					fmt.Println("Could not add to database", err)
					continue
				}
				tempQuery = ""
			}
			counter++
		}

		if tempQuery != "" {
			queryVal := "(" + rec[0] + "," + "'" + rec[1] + "'" + "," + "'" + rec[2] + "'" + "," + rec[3] + "," + rec[4] + "," + rec[5] + "," + "'" + rec[6] + "'" + "," + rec[7] + "," + rec[8] + "," + rec[9] + "," + rec[10] + "," + rec[11] + "," + rec[12] + "," + rec[13] + "," + rec[14] + "," + rec[15] + "," + rec[16] + ")"
			tempQuery = tempQuery + "," + queryVal
			tempQuery = query + tempQuery
			_, err2 := db.Exec(tempQuery)
			if err2 != nil {
				fmt.Println("Could not add to database", err)
				continue
			}
		}
	}
}

func main() {
	var err error
	var wg sync.WaitGroup
	workers := 18
	wg.Add(workers)

	// channel to pass the files
	chan1 := make(chan fs.FileInfo)

	// connect to sql server
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/NewYorkYellowTaxiData")
	if err != nil {
		fmt.Println("error connecting to sql server", err)
		return
	}
	defer db.Close()

	for i := 0; i < workers; i++ {
		// 18 workers for 18 files = 1 worker per file
		// the workers need to be ready before the files are opened
		go ReadFile(&wg, chan1, db)
	}

	// The directory of the file to be opened
	files1, err := os.Open(directory)
	if err != nil {
		fmt.Println("error opening directory:", err)
		return
	}
	defer files1.Close()

	//read the files from the directory
	fileInfos, err := files1.Readdir(-1)
	if err != nil {
		fmt.Println("error reading directory:", err)
		return
	}
	for _, fileInfo := range fileInfos {
		chan1 <- fileInfo
	}

	close(chan1)
	wg.Wait()
	fmt.Println("Over 100 million records inserted into the sql database!")
}
