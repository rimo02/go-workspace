package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

const Url = "http://xkcd.com"

type Result struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	Safe_title string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func fetch(n int) (*Result, error) {
	client := &http.Client{
		Timeout: 3 * time.Minute,
	}

	url := strings.Join([]string{Url, fmt.Sprintf("%d", n), "info.0.json"}, "/")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("http Request: %v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http err: %v", err)
	}

	var data Result

	if resp.StatusCode != http.StatusOK {
		data = Result{}
	} else {
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return nil, fmt.Errorf("json error:%v", err)
		}
	}
	resp.Body.Close()
	return &data, nil
}

type Job struct {
	number int
}

var jobs = make(chan Job, 100)
var results = make(chan Result, 100)
var resultcollection []Result

func allocateJobs(noOfJobs int) {
	for i := 0; i <= noOfJobs; i++ {
		jobs <- Job{i + 1}
	}
	close(jobs)
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		result, err := fetch(job.number)
		if err != nil {
			log.Printf("error in fetching: %v\n", err)
		}
		results <- *result
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i <= noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func getResults(done chan bool) {
	for result := range results {
		if result.Num != 0 {
			fmt.Printf("Retrieving issue #%d\n", result.Num)
			resultcollection = append(resultcollection, result)
		}
	}
	done <- true
}

func main() {
	noOfJobs := 3000
	go allocateJobs(noOfJobs)

	done := make(chan bool)
	go getResults(done)

	noOfWorkers := 100
	createWorkerPool(noOfWorkers)
	<-done

	data, err := json.MarshalIndent(resultcollection, "", "	")
	if err != nil {
		log.Fatal("json err: ", err)
	}
	err = writeToFile(data)
	if err != nil {
		log.Fatal(err)
	}

}

func writeToFile(data []byte) error {
	f, err := os.Create("xkcd.json")
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}
