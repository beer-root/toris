// This is the entry point of the TORIS application
package main

import (
  "fmt"
  "log"
  . "./toris/handlers"
  "code.google.com/p/couch-go"
  "code.google.com/p/goconf/conf"
  "code.google.com/p/goweb/goweb"
)

func main() {
  c, err := conf.ReadConfigFile("../conf/toris.conf")
  if err != nil {
    log.Fatal(err)
  }

  dbHost, _ := c.GetString("database", "host")
  dbPort, _ := c.GetInt("database", "port")
  dbName, _ := c.GetString("database", "name")

  serverPort, _ := c.GetInt("server", "port")

  // The CouchDB connection
  dbUrl := fmt.Sprintf("http://%s:%d/%s", dbHost, dbPort, dbName)
  log.Println("Connecting to CouchDB: " + dbUrl)
  db, err := couch.NewDatabaseByURL(dbUrl)
  if err != nil {
    log.Fatal(err)
  }

  // The REST controllers
  otterController := new(OtterAPIController)
  otterController.Db = db
  goweb.MapRest("/otter", otterController)

  leecherController := new(LeecherAPIController)
  leecherController.Db = db
  goweb.MapRest("/leecher", leecherController)

  // Add default result formatter (JSON)
  goweb.ConfigureDefaultFormatters()

  serverUrl := fmt.Sprintf(":%d", serverPort)
  log.Println("Starting server on: " + serverUrl)
  log.Fatal(goweb.ListenAndServe(serverUrl))

}
