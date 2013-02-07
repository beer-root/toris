// This is the entry point of the TORIS application
package main

import (
  "fmt"
  "log"
  "os"
  "./toris"
  "github.com/stretchrcom/goweb/goweb"
)


func main() {

  if toris.ShowModules {
    log.Println("Installed modules", toris.Context.ModuleList())
    os.Exit(0)
  }

  c := toris.Context.ConfigFile

  serverPort, _ := c.GetInt("server", "port")

  // Add default result formatter (JSON)
  goweb.ConfigureDefaultFormatters()

  // start the modules
  for name, _ := range(toris.Context.ModuleList()) {
    log.Println("Starting module " + name)
    err := toris.Context.Start(name)
    if err != nil {
      log.Println("[ERROR]", "An exception was raised", err)
    } else {
      log.Println("OK")
    }
  }

  serverUrl := fmt.Sprintf(":%d", serverPort)
  log.Println("Starting server on: " + serverUrl)
  log.Fatal(goweb.ListenAndServe(serverUrl))

}
