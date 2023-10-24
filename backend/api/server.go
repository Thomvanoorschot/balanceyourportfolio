package api

import (
	"fmt"

	"etfinsight/api/http"
)

func ListenAndServe(host string, port string) error {
	return http.ListenAndServe(fmt.Sprintf("%s:%s", host, port))
}
