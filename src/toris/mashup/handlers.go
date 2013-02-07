package mashup

import (
  "net/http"
  "github.com/stretchrcom/goweb/goweb"
)

// RESTful API Controller type for Otter
type OtterAPIController struct {
}

// Handle POST requests
func (cr *OtterAPIController) Create(cx *goweb.Context) {
  otter := NewOtter()
  // TODO rewrite using a tohva stored in the session
  // cr.Db.Insert(otter)
  cx.RespondWithData(otter)
}

// Handle DELETE requests
func (cr *OtterAPIController) Delete(id string, cx *goweb.Context) {
  //var otter Otter
  // TODO rewrite using a tohva stored in the session
  // currentRev, err := cr.Db.Retrieve(id, &otter)
  var err error = nil
  if err != nil {
  //  cr.Db.Delete(id, currentRev)
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
  // TODO rewrite using a tohva stored in the session
  var err error = nil
  //_, err := cr.Db.Retrieve(id, &otter)
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
