package port

import (
	"os"
	"strings"
)

func getPort() string {
	return strings.TrimSpace(os.Getenv("PORT"))
}
