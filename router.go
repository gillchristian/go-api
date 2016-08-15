package main

import (
  "net/http"

  "github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
  router := mux.NewRouter().StrictSlash(true)

  for _, route := range routes {
    var handler http.Handler
    handler = route.HandlerFunc
    handler = Logger(handler, route.Name)

    router.
      Methods(route.Method).
      Path(route.Pattern).
      Name(route.Name).
      Handler(handler)
  }

  router = setStaticHandler(router, "static")
  return router
}

func setStaticHandler (router *mux.Router, dir string) *mux.Router {
  staticHandler := http.StripPrefix(
    "/" + dir,
    http.FileServer(http.Dir("./" + dir)),
  )
  router.PathPrefix("/" + dir + "/").Handler(staticHandler)
  return router
}
