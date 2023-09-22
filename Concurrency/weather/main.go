package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const api = "8cc016291e702549194990020bf65a3d"

var data struct {
	Main struct {
		Temp float64 `json:"temp`
	} `json:"main"`
}

func fetch(city string, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, api)
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error finding data in city: %s", city)
		wg.Done()
		return
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		fmt.Printf("Error finding data in city: %s", city)
		wg.Done()
		return
	}
	ch <- fmt.Sprintf("Weather for city %s: %.2f", city, data.Main.Temp)
}

func main() {
	now := time.Now()
	cities := []string{"Toronto", "Delhi", "Paris", "Tokyo", "Sydney", "Barcelona"}
	ch := make(chan string)
	var wg sync.WaitGroup
	for _, city := range cities {
		wg.Add(1)
		go fetch(city, &wg, ch)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for msg := range ch {
		fmt.Println(msg)
	}
	fmt.Println("Time elapsed =", time.Since(now))
}
