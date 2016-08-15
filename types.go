package main

type jsonErr struct {
  Code int    `json:"code"`
  Text string `json:"text"`
}

type apiHome struct {
  Message string   `json:"message"`
  Routes  []string `json:"routes"`
}
