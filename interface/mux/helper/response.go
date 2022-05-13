package helper

import (
	"encoding/json"
	appError "github.com/i-jonathan/pharmacy-api/error"
	"log"
	"net/http"
)

func ReturnFailure(w http.ResponseWriter, err error) {
	switch err {
	case appError.BadRequest:
		w.WriteHeader(appError.BadRequest.Status)
		err2 := json.NewEncoder(w).Encode(appError.BadRequest.Response())
		log.Println(err2)
	case appError.NotFound:
		w.WriteHeader(appError.NotFound.Status)
		err2 := json.NewEncoder(w).Encode(appError.NotFound.Response())
		log.Println(err2)
	case appError.ServerError:
		w.WriteHeader(appError.ServerError.Status)
		err2 := json.NewEncoder(w).Encode(appError.ServerError.Response())
		log.Println(err2)
	}
}

func ReturnSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if msg, ok := data.(map[string]string); ok {
		err := json.NewEncoder(w).Encode(msg)
		if err != nil {
			log.Println(err)
		}
	} else {
		resp := struct {
			Message string      `json:"message"`
			Data    interface{} `json:"data"`
		}{"Success", data}
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Println(err)
		}
	}
}

func ReturnDelete(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}