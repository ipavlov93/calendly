package repository

import "fmt"

var ErrDuplicate = fmt.Errorf("db: duplicate error")
var ErrNoRows = fmt.Errorf("db: no rows in result set")
