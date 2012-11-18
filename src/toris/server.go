package toris

import (
  "code.google.com/p/couch-go"
  "code.google.com/p/goweb/goweb"
  "net/http"
)

// RESTful API Controller type
type OtterAPIController struct {
  Db couch.Database
}

// Handle POST requests
func (cr *OtterAPIController) Create(cx *goweb.Context) {
  otter := NewOtter()
  cr.Db.Insert(otter)
  cx.RespondWithData(otter)
}

// Handle DELETE requests
func (cr *OtterAPIController) Delete(id string, cx *goweb.Context) {
  var otter Otter
  currentRev, err := cr.Db.Retrieve(id, &otter)
  if err != nil {
    cr.Db.Delete(id, currentRev)
    cx.RespondWithOK()
  } else {
    cx.RespondWithNotFound()
  }
}
func (cr *OtterAPIController) DeleteMany(cx *goweb.Context) {
  cx.RespondWithStatus(http.StatusForbidden)
}

// Handle GET requests
func (cr *OtterAPIController) Read(id string, cx *goweb.Context) {
  var otter Otter
  _, err := cr.Db.Retrieve(id, &otter)
  if err != nil {
    cx.RespondWithData(otter)
  } else {
    cx.RespondWithNotFound()
  }
}
func (cr *OtterAPIController) ReadMany(cx *goweb.Context) {
  cx.RespondWithStatus(http.StatusForbidden)
}

// Handle PUT requests
func (cr *OtterAPIController) Update(id string, cx *goweb.Context) {
  cx.RespondWithStatus(http.StatusForbidden)
}
func (cr *OtterAPIController) UpdateMany(cx *goweb.Context) {
  cx.RespondWithStatus(http.StatusForbidden)
}
