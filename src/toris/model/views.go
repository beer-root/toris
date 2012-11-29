package toris

import (
  "code.google.com/p/couch-go"
)

type Design struct {
  Id string `json: "_id"`
  Rev string `json: "_rev,omitempty"`
  Language string `json:"language"`
  Views map[string] View `json:"views,omitempty"`
  ValidateDocUpdate string `json: "validate_doc_update,omitempty"`
}

type View struct {
  Map string `json:"map"`
  Reduce string `json:"map,omitempty"`
}

type Database struct {
  couch.Database
}

func (d Database) SaveDesign(design Design) (string, string, error) {
  return d.Insert(design)
}
