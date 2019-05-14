package walk

import (
	_ "golang.org/x/crypto/acme"
	_ "github.com/go-sql-driver/mysql"
	_ "go.opencensus.io/tag"
)

// BestRoute tells you the best dog walking route in your neighborhood.
func BestRoute() string {
	return "the best place to walk your dog is by your side"
}
