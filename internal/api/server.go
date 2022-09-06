package api

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/flexoid/translators-map-go/ent"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Server struct {
	EntDB  *ent.Client
	Logger *zap.SugaredLogger
}

func (s *Server) Start(bindAddr string) error {
	router := s.setupRouter()

	err := http.ListenAndServe(bindAddr, router)
	return errors.Wrap(err, "unable to start the server")
}

func (s *Server) setupRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Group(func(apiRouter chi.Router) {
		apiRouter.Use(middleware.RealIP)
		apiRouter.Use(middleware.Recoverer)
		apiRouter.Use(middleware.SetHeader("Content-Type", "application/json"))

		translatorsController := &TranslatorController{Server: s}
		apiRouter.Get("/api/translators", translatorsController.GetTranslators)

		configController := &ConfigController{Server: s}
		apiRouter.Get("/api/config", configController.GetConfig)

		languagesController := &LanguageController{Server: s}
		apiRouter.Get("/api/languages", languagesController.GetLanguages)
	})

	s.setupFileServer(router, "web")

	return router
}

// Serve static files of frontend app.
func (s *Server) setupFileServer(router chi.Router, path string) {
	workDir, _ := os.Getwd()
	root := http.Dir(filepath.Join(workDir, path))

	fileServer := http.FileServer(root)

	router.
		Get("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Cache-Control", "must-revalidate")
			fileServer.ServeHTTP(w, r)
		}))
}
