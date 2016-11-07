// Tutorial - on net/http
package main

import (
	//"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", r.URL.Path[1:])
	//w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//myItems := []string{"item1", "item2", "item3"}
	//a, _ := json.Marshal(myItems)

	//w.Write(a)
	//return
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Foo")
	http.ListenAndServe(":8081", nil)
	//log.Println(http.ListenAndServe(":8081", nil))
}
