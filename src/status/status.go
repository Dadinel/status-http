package status

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func getRandStatus() int {
	rand.NewSource(time.Now().UnixNano())
	status := (rand.Intn(4) + 2) * 100
	return status
}

func getURLStatus(r *http.Request) int {
	paramCode := strings.Replace(r.URL.Path, "/", "", -1)
	status, _ := strconv.Atoi(paramCode)
	return status
}
