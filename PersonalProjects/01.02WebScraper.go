package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func ReadWriteFile(wg *sync.WaitGroup, chan1 chan string) {
	// take the websites form the channel
	defer wg.Done()
	if err := recover(); err != nil {
		fmt.Println(err)
	}
	for {
		rec2, ok := <-chan1
		if !ok {
			fmt.Println("Channel closed")
		}

		// make an http call to the website
		client := &http.Client{
			Timeout: 30 * time.Second,
		}
		rec2 = "http://" + rec2
		req, _ := http.NewRequest("GET", rec2, nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/115.0")
		resp, err := client.Do(req)
		if os.IsTimeout(err) {
			log.Printf("timeout error: %v\n", err)
			continue
		}
		if err != nil {
			fmt.Println("Could not connect to Website", rec2, err)
			continue
		}

		// Read the response
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		str1 := bytes.NewBuffer(body).String()
		if err != nil {
			fmt.Println("Could not read response body", rec2, err)
			continue
		}

		//Create a file with the website's name
		t := strings.Replace(rec2, ":", "", -1)
		t = strings.Replace(t, "/", "_", -1)
		path := "C:\\Users\\tanma\\OneDrive\\Documents\\1-mWebsites\\" + string(t[6]) + string(t[7]) + "\\"
		if err := os.Mkdir(path, os.ModePerm); os.IsNotExist(err) {
			err3 := os.Mkdir(path, os.ModePerm)
			if err3 != nil {
				fmt.Println("Could not make file with the same initial letter as the website", err3)
				continue
			}
		}
		name := path + t + ".txt"
		f, err := os.Create(name)
		if err != nil {
			fmt.Println("Could not create file", name, err)
			continue
		}

		//Write the response to a file
		_, err2 := f.WriteString(str1)
		if err2 != nil {
			fmt.Println("Could not write to file", name, err2)
			continue
		}
		defer f.Close()
		d := time.Microsecond
		time.Sleep(d)
	}
}

func File(wg *sync.WaitGroup, workers int) {
	chan1 := make(chan string)

	for i := 0; i < workers; i++ {
		go ReadWriteFile(wg, chan1)
	}

	//Read from the file with 1 million websites
	filePath := "C:\\Users\\tanma\\Downloads\\top-1m.csv"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error while opening the file", err)
		fmt.Println(err)
		return
	}
	reader := csv.NewReader(file)
	reader.Comma = ','
	if err != nil {
		fmt.Println("Error reading records", err)
		return
	}

	//Pass the websites to the channel
	for {
		rec, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Could not read the websites file")
			log.Fatal(err)
			return
		}
		chan1 <- rec[1]
	}
	close(chan1)
}

func main() {
	var wg sync.WaitGroup
	totalWorkers := 100
	wg.Add(totalWorkers)
	go File(&wg, totalWorkers)
	wg.Wait()
}
