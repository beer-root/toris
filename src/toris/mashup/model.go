package mashup

import (
  "github.com/beer-root/tohva"
  "time"
)


// An Otter represents the image and its associated meta-data
type Otter struct {
  tohva.WithIdRev
  Path string // path to the otter file
  Tags []string // the tags of this otter
  Date time.Time // when published
  Votes map[string] int // leecher_id -> score
}

// Create an empty Otter
func NewOtter() Otter {
  return Otter{Date: time.Now()}
}

// views on otters

