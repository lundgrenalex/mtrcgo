package api

import "net/http"

func GetStatus(w http.ResponseWriter, r *http.Request) {
    if (r.URL.Path != "/") {
        SendResponse(HttpResponse{404, "Not Found!"}, w)
        return
    }
}
