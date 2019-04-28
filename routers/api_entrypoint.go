package routers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tflac_cw/kernel"
	"tflac_cw/models"

	"github.com/gorilla/mux"
)

/*EntryPoint - структура для точки входа в апи*/
type EntryPoint struct {
	KernelLanguage *kernel.Kernel
}

/*Init - инциализация*/
func (entry *EntryPoint) Init() *EntryPoint {
	ker := &kernel.Kernel{}
	entry.KernelLanguage = ker.New()
	return entry
}

func (entry *EntryPoint) responseWithJSON(w http.ResponseWriter, r *http.Request, httpstatus int, payload interface{}) {
	if r.Header.Get("Content-Type") != "application/json" {
		log.Println("Type in not format!")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)

		response, _ := json.Marshal(map[string]string{"error": "you packet in non json format!"})
		w.Write(response)
	} else {
		response, _ := json.Marshal(payload)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpstatus)

		w.Write(response)
	}
}

/*HandledNewTaskForAnalys - обработчик задания для анализов*/
func (entry *EntryPoint) HandledNewTaskForAnalys(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var model models.EntryPayload
	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		entry.responseWithJSON(w, r, http.StatusOK, map[string]string{"error": "can not recognize input payload"})
	}
	entry.KernelLanguage.Input(model.Text)
	response := models.ResponsePayload{
		LexerAnalysPart: models.LexerPart{
			Tokens: entry.KernelLanguage.LexerAnalys.GetTokens(),
			Errors: entry.KernelLanguage.LexerAnalys.GetErrors(),
		},
		SyntaxAnalysPart: models.SyntaxPart{
			Errors: entry.KernelLanguage.SyntaxAnalys.GetErrors(),
		},
	}

	fmt.Println("Result: ", response)

	entry.responseWithJSON(w, r, http.StatusOK, response)
}

/*CreateEntryPoint - создание точки входа в центральное апи для начала анализа лекс + синт.*/
func (entry *EntryPoint) CreateEntryPoint() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/cw", entry.HandledNewTaskForAnalys).Methods("POST")

	return router
}
