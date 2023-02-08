package helper

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, status int, record interface{}, err error) {
	response := []byte{}
	writeBody := false

	if err != nil {
		writeBody = true
		response, _ = json.Marshal(map[string]string{"error_description": err.Error()})
	} else if record != nil {
		writeBody = true
		response, _ = json.Marshal(record)
	}

	w.WriteHeader(status)
	if writeBody {
		w.Write(response)
	}
}
