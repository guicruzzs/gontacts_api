package main

import(
  "fmt"
  "net/http"
  "./auth"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("GET", r.URL)
  fmt.Fprintf(w, "GET %s", r.URL.Path[1:])
  cas.Blink()
}

func main(){
  http.HandleFunc("/", handler)
  http.ListenAndServe(":3000", nil)
}
