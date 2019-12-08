/*
 * Hyperion
 * Étienne "Maiste" Marais
 */

package api

import (
  "fmt"
  "log"
  "net/http"               // Http Server
  "github.com/gorilla/mux" // Routing
)

/* Launching function */
func Display() {
  fmt.Println("Hyperion Template server v0.1")
}

/* Test connection with the server */
func handleTest(w http.ResponseWriter, r* http.Request) {
  fmt.Println("[TEST] Connection")
  w.WriteHeader(http.StatusOK)
}

/* Handle template download */
func handleTemplate(w http.ResponseWriter, r* http.Request) {
  request := mux.Vars(r)
  template := "templates/" + request["template"]
  fmt.Println("[REQUEST] Download", template)
  rawTemplate, err := ImportJson(template)
  if err != nil {
    fmt.Println("[ERROR]", err)
    w.WriteHeader(http.StatusNotFound)
  } else {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(rawTemplate)
  }
}

/* Handle connextions on port 8080 */
func ConnectionHandler() {
  route := mux.NewRouter().StrictSlash(true)
  route.HandleFunc("/", handleTest)
  route.HandleFunc("/{template}", handleTemplate)
  log.Fatal(http.ListenAndServe(":8080", route))
}
