package main

import (
  "./toris"
  "./toris/mashup"
  "code.google.com/p/goconf/conf"
  "github.com/stretchrcom/goweb/goweb"
)

type MashupModule struct {}

func (m MashupModule) Name() string {
  return "MashupModule"
}

func (m MashupModule) Start(config conf.ConfigFile) error {
  // TODO check for DB, etc...
  controller := mashup.OtterAPIController{}
  // bind the controller
  goweb.MapRest("/otter", controller)
  return nil
}

func (m MashupModule) Stop() error {
  return nil
}

func init() {
  toris.Context.Register(MashupModule{})
}
