package Middleware

import (
	"encoding/json"
	"net/http"
	"time"
)

/*Responser - тип содержащий базовые методы для обработок ошибочных и успешных запросов*/
type Responser struct {
	Tururu string
}

/*ResponseWithError - ответ с ошибкой*/
func (resp *Responser) ResponseWithError(w http.ResponseWriter, r *http.Request, httpStatus int, payload interface{}, timestart time.Time) {
	resp.ResponseWithJSON(w, r, httpStatus, payload, timestart)
}

/*ResponseWithJSON - ответ в формате json*/
func (resp *Responser) ResponseWithJSON(w http.ResponseWriter, r *http.Request, httpStatus int, payload interface{}, timestart time.Time) {
	log := Logger{}
	log.ConfigTrace(r.Method, r.RequestURI, r.Header.Get("Content-Type"), time.Since(timestart)).PrintTraceRoute()

	if r.Header.Get("Content-Type") != "application/json" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		response, _ := json.Marshal(map[string]string{"error": "you packet in non json format!"})
		w.Write(response)
	} else {
		response, _ := json.Marshal(payload)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpStatus)
		w.Write(response)
	}
}
