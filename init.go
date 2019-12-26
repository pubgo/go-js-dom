package godom

import (
	"fmt"
	"github.com/pubgo/godom/version"
	"os"
)

func init() {
	if os.Getenv("version") != "" {
		fmt.Printf("godom Version %s, BuildV %s, CommitV %s\n", version.Version, version.BuildV, version.CommitV)
	}
}
