package status_controller

import (
	"fmt"
	"net/http"
)

func Status(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := `{"status": "ok"}`
	_, err := w.Write([]byte(data))
	if err != nil {
		fmt.Println("Error Status Controller: " + err.Error())
	}
}