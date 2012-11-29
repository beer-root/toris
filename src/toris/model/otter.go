package toris

import (
  "time"
)


// An Otter represents the image and its associated meta-data
type Otter struct {
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

