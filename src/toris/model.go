package toris

import "github.com/beer-root/tohva"

type User struct {
  tohva.WithIdRev
  Name string `json:"name"`
  FirstName string `json:"firstname"`
  LastName string `json:"lastname"`
  Email string `json:"email_address"`
}
