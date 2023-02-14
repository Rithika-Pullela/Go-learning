package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
	"io/ioutil"
)

type User struct{
	Name string
	Job string
	Id string
	CreatedAt string
}

func main() {

    values := User{Name: "Rithika", Job: "sd"}
	m:=User{}
    json_data, err := json.Marshal(values)

    if err != nil {
        log.Fatal(err)
    }

    resp, err := http.Post("https://reqres.in/api/users", "json",
        bytes.NewBuffer(json_data))

    if err != nil {
        log.Fatal(err)
    }
	defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }
    err1 := json.Unmarshal(body, &m)
    if err != nil {
        log.Fatalln(err1)
    }
    fmt.Println(m.Id)
    fmt.Println(m.CreatedAt)
}


