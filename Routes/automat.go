package Routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gitlab.com/SmokeMazero/TFLaC/lab2/Middleware"
	"gitlab.com/SmokeMazero/TFLaC/lab2/Models"
)

/*AutomatRoute - структура для роутера*/
type AutomatRoute struct {
	AmountExceptions int32
}

/*CreateNewTextAnalys - функция для создания анализа отправляемого текста с фронта*/
func (setRoute *AutomatRoute) CreateNewTextAnalys(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	responser := Middleware.Responser{}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		responser.ResponseWithError(w, r, http.StatusBadRequest, map[string]string{
			"error": "not handled models in your payload",
		}, startTime)
		return
	}

	var payload Models.NewCodePayload
	err = json.Unmarshal(b, &payload)

	if err != nil {
		responser.ResponseWithError(w, r, http.StatusBadRequest, map[string]string{
			"error": "not have unmarshled your payload!",
		}, startTime)
		return
	}

	responser.ResponseWithJSON(w, r, http.StatusOK, map[string]interface{}{
		"data":            payload.Value,
		"errors":          0,
		"warnings":        0,
		"errorMessages":   []string{""},
		"warningMessages": []string{""},
	}, startTime)
}

/*SettingRouter - конфигурация маршрута на роутере automat*/
func (setRoute *AutomatRoute) SettingRouter(router *mux.Router) {
	router.HandleFunc("/v0.1/2lab", setRoute.CreateNewTextAnalys).Methods("POST")
	return
}
