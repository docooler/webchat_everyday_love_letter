package main 

import (
         "io"
         "net/http"
         "log"
         "fmt"
         )

func hellohandler( w http.ResponseWriter, r * http.Request) {
   io.WriteString(w, "Hello, world!")
}

func main() {
   http.HandleFunc("/hello", hellohandler)
   err := http.ListenAndServe(":8080", nil)
   if err != nil {
      log.Fatal("God like")
   }
   fmt.Println("run server at 8080")
}