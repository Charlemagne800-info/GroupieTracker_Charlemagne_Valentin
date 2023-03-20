package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
	"time"
)

var day_Joke = template.Must(template.ParseFiles("templates/Day_Joke.html"))
var random = template.Must(template.ParseFiles("templates/Random_Joke.html"))
var index = template.Must(template.ParseFiles("templates/index.html"))

type Joke_Day struct {
	Response string `json:"response"`
	Joke     struct {
		Question string `json:"question"`
		Answer   string `json:"answer"`
	} `json:"joke"`
}

type Joke_Random struct {
	Response_Random string `json:"response"`
	Joke            struct {
		Question_Random string `json:"question"`
		Answer_Random   string `json:"answer"`
	} `json:"joke"`
}

type Joke_Index struct {
	Response_Index string `json:"response"`
	Joke           struct {
		Question_Index string `json:"question"`
		Answer_Index   string `json:"answer"`
	} `json:"joke"`
}

func main() {
	styleServer := http.FileServer(http.Dir("assets"))
	styleServer2 := http.FileServer(http.Dir("css"))
	styleServer3 := http.FileServer(http.Dir("js"))
	http.Handle("/css/", http.StripPrefix("/css/", styleServer2))
	http.Handle("/js/", http.StripPrefix("/js/", styleServer3))
	http.Handle("/assets/", http.StripPrefix("/assets/", styleServer))
	//fmt.Println(styleServer)
	http.HandleFunc("/", handler_Index)
	http.HandleFunc("/Day_Joke.html", handler_Day)
	http.HandleFunc("/Random_Joke.html", handler_Random)
	fmt.Printf("Serveur lancé sur le port 9100 : http://localhost:9100")
	http.ListenAndServe(":9100", nil)
}

func handler_Index(w http.ResponseWriter, r *http.Request) {
	url := "https://blague.xyz/api/joke/random"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	req.Header.Set("User-Agent", "spacecount-total")

	timeClient := http.Client{
		Timeout: time.Second * 2,
	}

	res, getErr := timeClient.Do(req)
	if getErr != nil {
		fmt.Println(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}

	var temp2 Joke_Index

	total := json.Unmarshal(body, &temp2)
	if total != nil {
		fmt.Println(temp2.Joke.Answer_Index)
	}

	data := Joke_Index{
		Response_Index: temp2.Response_Index,
		Joke: struct {
			Question_Index string `json:"question"`
			Answer_Index   string `json:"answer"`
		}{
			Question_Index: temp2.Joke.Question_Index,
			Answer_Index:   temp2.Joke.Answer_Index,
		},
	}

	index.Execute(w, data)
}

func handler_Day(w http.ResponseWriter, r *http.Request) {
	url := "https://blague.xyz/api/joke/day"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	req.Header.Set("User-Agent", "spacecount-total")

	timeClient := http.Client{
		Timeout: time.Second * 2,
	}

	res, getErr := timeClient.Do(req)
	if getErr != nil {
		fmt.Println(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}

	var temp2 Joke_Day

	total := json.Unmarshal(body, &temp2)
	if total != nil {
		fmt.Println(temp2.Joke.Answer)
	}

	data := Joke_Day{
		Response: fmt.Sprintf(temp2.Response),
		Joke: struct {
			Question string `json:"question"`
			Answer   string `json:"answer"`
		}{
			Question: fmt.Sprintf(temp2.Joke.Question),
			Answer:   fmt.Sprintf(temp2.Joke.Answer),
		},
	}

	day_Joke.Execute(w, data)
}

func handler_Random(w http.ResponseWriter, r *http.Request) {
	url := "https://blague.xyz/api/joke/random"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	req.Header.Set("User-Agent", "spacecount-total")

	timeClient := http.Client{
		Timeout: time.Second * 2,
	}

	res, getErr := timeClient.Do(req)
	if getErr != nil {
		fmt.Println(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}

	var temp2 Joke_Random

	total := json.Unmarshal(body, &temp2)
	if total != nil {
		fmt.Println(temp2.Joke.Answer_Random)
	}

	data := Joke_Random{
		Response_Random: fmt.Sprintf(temp2.Response_Random),
		Joke: struct {
			Question_Random string `json:"question"`
			Answer_Random   string `json:"answer"`
		}{
			Question_Random: fmt.Sprintf(temp2.Joke.Question_Random),
			Answer_Random:   fmt.Sprintf(temp2.Joke.Answer_Random),
		},
	}

	random.Execute(w, data)
}
