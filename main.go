package main

import (
  "fmt"
  "github.com/CiscoZeus/go-zeusclient"
  "net/http"
 "os"
 "runtime"
  "encoding/json"
  //"io/ioutil"
  // "strconv"

 )

type Log map[string]interface{}

// A collection of logs.
type LogList struct {
        Name string
        Logs []Log
}

func (lst LogList) MarshalJSON() ([]byte, error) {
        return json.Marshal(lst.Logs)
}


 func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "This is a verification test %s!", r.URL.Path[1:])
  //fmt.Fprintf(w, "Hello there, this was passed in from the URL :%s\n", r.URL.Path[1:])
 }

 func main() {
 // print env variables
 fmt.Println("OS:", runtime.GOOS)
 fmt.Println("USER:", os.Getenv("USER"))
 fmt.Println("HOME:", os.Getenv("HOME"))
 fmt.Println("Token:", os.Getenv("ZEUS_TOKEN"))
  hostname, err := os.Hostname()
    if err != nil{}
fmt.Println("HOSTNAME:", hostname)

 // send env variables to zeus
z := &zeus.Zeus{ApiServ: "http://api.ciscozeus.io", Token: os.Getenv("ZEUS_TOKEN")}
logs := zeus.LogList{
 Name: "Eliza-Khorashahi",
 Logs: []zeus.Log{
 zeus.Log{
 "OS":       runtime.GOOS,
 "USER":     os.Getenv("USER"),
 "HOME":     os.Getenv("HOME"),
 "HOSTNAME": hostname},
 },
 }
 
 suc, err := z.PostLogs(logs)
 if err != nil {
   fmt.Println("Zeus Error response:", err.Error())
 } else {
  fmt.Println("Zeus Success response:", suc)
 }

  fmt.Printf("Starting http server...\n")
  http.HandleFunc("/", handler)
 err2 := http.ListenAndServe(":8080", nil)
 fmt.Println(err2.Error())

}