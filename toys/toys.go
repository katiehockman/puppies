package toys

import (
       "fmt"
	_ "golang.org/x/crypto/acme"
	_ "github.com/go-sql-driver/mysql"
	_ "go.opencensus.io/tag"
)

// Reminder reminds you to buy your dog a toy.
func Reminder() {
	fmt.Print("go buy a toy for your dog!!")
}
