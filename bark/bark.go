package bark

import (
	_ "golang.org/x/crypto/acme"
	_ "github.com/go-sql-driver/mysql"
	_ "go.opencensus.io/tag"
)

// Meaning tells you want your dog's bark means.
func Meaning() string {
	return "your dog loves you <3"
}
