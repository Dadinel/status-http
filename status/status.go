package status

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// GetRandStatus : Retorna um status http aleatório
func GetRandStatus() int {
	rand.NewSource(time.Now().UnixNano())
	status := (rand.Intn(4) + 2) * 100
	return status
}

// GetURLStatus : Retorna o status http conforme recebido na url
func GetURLStatus(r *http.Request) int {
	paramCode := strings.Replace(r.URL.Path, "/", "", -1)
	status, _ := strconv.Atoi(paramCode)
	return status
}
