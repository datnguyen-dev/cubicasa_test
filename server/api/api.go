package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strings"

	"datnguyen.cubicasa.test/store"
	"github.com/go-chi/chi"
)

//Handler - Handle router
type Handler struct {
	inspector *inspector
	process   *process
	router    chi.Router
	store     store.Store
}

//New - init API
func New(store store.Store) *Handler {
	h := &Handler{}
	h.store = store
	h.inspector = &inspector{h}
	h.process = &process{h}
	h.initRouter()
	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *Handler) urlParam(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}

func (h *Handler) parseRequest(r *http.Request, data interface{}) error {
	const maxRequestLen = 16 * 1024 * 1024
	lr := io.LimitReader(r.Body, maxRequestLen)
	return json.NewDecoder(lr).Decode(data)
}

func (h *Handler) render(w http.ResponseWriter, status int, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		h.logError("marshal json: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)
}

func (h *Handler) renderError(w http.ResponseWriter, status int, code, message string) {
	response := struct {
		Error struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}{}
	response.Error.Code = code
	response.Error.Message = message
	h.render(w, status, response)
}

func (h *Handler) logError(format string, a ...interface{}) {
	pc, _, _, _ := runtime.Caller(1)
	callerNameSplit := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	funcName := callerNameSplit[len(callerNameSplit)-1]
	fmt.Printf("%s: %s", funcName, fmt.Sprintf(format, a...))
}
