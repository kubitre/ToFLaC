package Routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*EntrypointRoute - роутер главный*/
type EntrypointRoute struct {
	AmountErrr int32
}

/*CreateAllRoutes - формирование главного роутера*/
func (setRoute *EntrypointRoute) CreateAllRoutes() {
	router := mux.NewRouter()
	autom := AutomatRoute{}
	autom.SettingRouter(router)

	log.Fatal(http.ListenAndServe(":9999", router))
}
