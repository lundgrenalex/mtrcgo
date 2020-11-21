package api

import (
	"github.com/lundgrenalex/mtrcgo/storage"
	"net/http"
)

func GetStatus(crud storage.CRUDStorage, w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		SendResponse(HttpResponse{404, "Not Found!"}, w)
		return
	}
	metrics := crud.GetAllRecords()

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(metrics.Marshall())
}
