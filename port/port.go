package port

import (
	"fmt"
	"os"
	"strings"
)

// GetPort : Retorna a porta da aplicação
func GetPort() (string, error) {
	port := strings.TrimSpace(os.Getenv("PORT"))

	if port == "" {
		return "", fmt.Errorf("Invalid port")
	}

	return port, nil
}
