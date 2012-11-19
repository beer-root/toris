package toris

import (
  "code.google.com/p/couch-go"
  "code.google.com/p/goweb/goweb"
)


// A Leccher represents a toris user
type Leecher struct {
  Username string
  Mail string
  Password string
}

// RESTful API Controller type for Leecher
type LeecherAPIController struct {
  Db couch.Database
}

// Handle POST requests
func (cr *LeecherAPIController) Create(cx *goweb.Context) {
  leecher := Leecher{}

  leecher.Username = cx.Request.FormValue("username")
  leecher.Mail = cx.Request.FormValue("mail")
  leecher.Password = cx.Request.FormValue("password")
  cr.Db.Insert(leecher)
  cx.RespondWithData(leecher)
}

// Handle DELETE requests
func (cr *LeecherAPIController) Delete(id string, cx *goweb.Context) {
  var leecher Leecher
  if currentRev, err := cr.Db.Retrieve(id, &leecher) ; err != nil {
    cx.RespondWithNotFound()
  } else {
    cr.Db.Delete(id, currentRev)
    cx.RespondWithOK()
  }
}

// Handle GET requests
func (cr *LeecherAPIController) Read(id string, cx *goweb.Context) {
  var leecher Leecher
  if _, err := cr.Db.Retrieve(id, &leecher) ; err != nil {
    cx.RespondWithNotFound()
  } else {
    cx.RespondWithData(leecher)
  }
}
