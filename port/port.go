package port

import (
	"os"
	"strings"
)

func GetPort() string {
	return strings.TrimSpace(os.Getenv("PORT"))
}
