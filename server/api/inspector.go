package api

import (
	"net/http"
)

type inspector struct {
	*Handler
}

type inspectInfo struct {
	AppName string `json:"name"`
	Version string `json:"version"`
}

func (h *inspector) getInspect(w http.ResponseWriter, r *http.Request) {
	data := inspectInfo{
		AppName: "NguyenTanDat-Cubicasa",
		Version: "1.0.000.1",
	}
	h.render(w, http.StatusOK, data)
}
