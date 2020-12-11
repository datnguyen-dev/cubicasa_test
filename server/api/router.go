package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (h *Handler) initRouter() {
	r := chi.NewRouter()

	// Use middleware: CORS
	r.Use(h.Allow)

	// Public routes
	r.Group(func(r chi.Router) {
		r.Get("/inspect", h.inspector.getInspect)

		r.Get("/checkdb", h.process.checkDB)
		r.Get("/install", h.process.install)

		r.Post("/users", h.process.createUser)
		r.Post("/hubs", h.process.createHub)
		r.Post("/teams", h.process.createTeam)

		r.Put("/users", h.process.updateUser)
		r.Put("/hubs", h.process.updateHub)
		r.Put("/teams", h.process.updateTeam)

		r.Delete("/users/{id}", h.process.deleteUser)
		r.Delete("/hubs/{id}", h.process.deleteHub)
		r.Delete("/teams/{id}", h.process.deleteTeam)

		r.Get("/searchhubs/{name}", h.process.searchHub)
		r.Get("/searchteams/{name}", h.process.searchTeam)

		r.Put("/jointeam/{teamid}/{hubid}", h.process.joinTeam)
		r.Put("/joinuser/{userid}/{teamid}/{role}", h.process.joinUser)
	})

	h.router = r
}

// Allow CORS
func (h *Handler) Allow(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if len(origin) == 0 {
			origin = "*"
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, HEAD, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Encoding, Authorization, Content-Length, Content-Type, X-SSO-Ticket")
		w.Header().Set("Access-Control-Expose-Headers", "Accept, Accept-Encoding, Authorization, Content-Length, Content-Type, X-SSO-Ticket")

		// If this was preflight options request let's write empty ok response and return
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			w.Write(nil)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
