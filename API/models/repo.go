package models

import (
  //"strings"
)

type Repo struct {
  ID string
  Level uint
  URL string
  Stargazers []User
}
