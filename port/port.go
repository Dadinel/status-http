package port

import (
	"os"
	"strings"
)

// GetPort : Retorna a porta da aplicação
func GetPort() string {
	return strings.TrimSpace(os.Getenv("PORT"))
}
