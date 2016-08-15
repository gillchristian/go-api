package main

import (
  "os"
  "fmt"
  "log"
  "net/http"
)

func main() {
  PORT := logAndGetPort()
  router := NewRouter()
  log.Fatal(http.ListenAndServe(":" + PORT, router))
}

func logAndGetPort() string {
  PORT := os.Getenv("PORT")
  if len(PORT) == 0 {
    PORT = "8080"
  }
  message := "Listening on port %v\nVisit http://localhost:%v\n"

  fmt.Printf(message, PORT, PORT)
  return PORT
}
