package toris

import "time"

// An Otter represents the image and its associated meta-data
type Otter struct {
  Path string
  Tags []string
  Date time.Time
}

// Create an empty Otter
func NewOtter() Otter {
  return Otter{Date: time.Now()}
}
