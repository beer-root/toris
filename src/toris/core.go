package toris

import (
  "flag"
  "log"
  "os"
  "code.google.com/p/goconf/conf"
  "github.com/beer-root/tohva"
  "github.com/gorilla/sessions"
)

// the Toris context accessible from the modules
var Context Toris

var ShowModules bool

func init() {

  // command line options
  var confFile string
  flag.StringVar(&confFile, "config", "/etc/toris/toris.conf", "The configuration file")

  flag.BoolVar(&ShowModules, "module-list", false, "Print the list of loaded modules and exits")

  flag.Parse()

  // read the configuration file
  config, err := conf.ReadConfigFile(confFile)
  if err != nil {
    log.Fatal(err)
    os.Exit(1)
  }

  // initializes the database connection
  dbHost, _ := config.GetString("database", "host")
  dbPort, _ := config.GetInt("database", "port")
  dbName, _ := config.GetString("database", "name")
  adminName, _ := config.GetString("database", "admin_name")
  adminPwd, _ := config.GetString("database", "admin_password")

  couch := tohva.CreateCouchClient(dbHost, dbPort)

  // the secret token for secure cookies
  secretToken, _ := config.GetString("sessions", "secret")

  Context = Toris {
    map[string]InstalledModule{},
    config,
    &couch,
    dbName,
    adminName,
    adminPwd,
    sessions.NewCookieStore([]byte(secretToken)) }

  // start all the modules
  for _, m := range(Context.modules) {
    log.Println("Starting module " + m.module.Name())
  }

}

// The Toris Instance

type Toris struct {
  modules map[string]InstalledModule
  ConfigFile *conf.ConfigFile
  Couch *tohva.CouchDB
  CorisDb string
  couchAdminName string
  couchAdminPwd string
  Sessions sessions.Store
}

// registers the module
func (t Toris) Register(module Module) {
  t.modules[module.Name()] = InstalledModule { module, Stopped }
}

// return the list of installed module names along with their status
func (t Toris) ModuleList() map[string]Status {
  result := map[string]Status{}
  for k, m := range(t.modules) {
    result[k] = m.status
  }
  return result
}

// starts a module by name
func (t Toris) Start(name string) error {
  return t.modules[name].module.Start(*t.ConfigFile)
}

// stops a module by name
func (t Toris) Stop(name string) {
  t.modules[name].module.Stop()
}

// The Modules

type InstalledModule struct {
  module Module
  status Status
}

type Status int
const(
  Stopped Status = iota
  Started
)

// a module can be started and stopped and has an id
type Module interface {
  Name() string
  Start(c conf.ConfigFile) error
  Stop() error
}
