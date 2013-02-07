package main

import (
  "./toris"
  "github.com/beer-root/tohva"
  "code.google.com/p/goconf/conf"
)

// the model
type Item struct {
  tohva.WithIdRev
  Type string `json:"type"`
  Name string
  Owner toris.User
}

type Loan struct {
  tohva.WithIdRev
  Item Item `json:"item"`
  To toris.User `json:"to"`
}

type LoansModule struct {}

func (m LoansModule) Name() string {
  return "LoansModule"
}

func (m LoansModule) Start(config conf.ConfigFile) error {
  // TODO check for DB, etc...
  return nil
}

func (m LoansModule) Stop() error {
  return nil
}

func init() {
  toris.Context.Register(LoansModule{})
}
