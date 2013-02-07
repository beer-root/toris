// This is the entry point of the TORIS application
package main

import (
  "fmt"
  "log"
  "./toris"
  "github.com/stretchrcom/goweb/goweb"
)


func main() {

  c := toris.Context.ConfigFile

  serverPort, _ := c.GetInt("server", "port")

  // Add default result formatter (JSON)
  goweb.ConfigureDefaultFormatters()

  serverUrl := fmt.Sprintf(":%d", serverPort)
  log.Println("Starting server on: " + serverUrl)
  log.Fatal(goweb.ListenAndServe(serverUrl))

}
