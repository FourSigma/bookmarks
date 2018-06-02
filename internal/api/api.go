package api

import (
	"net/http"

	"github.com/FourSigma/bookmarks/internal/service"
	"github.com/FourSigma/bookmarks/pkg/log"
	"github.com/FourSigma/bookmarks/pkg/opengraph"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Resource interface {
	Routes(chi.Router)
}

func InitalizeAPIResources(rs ...Resource) (r *chi.Mux) {
	r = chi.NewRouter()
	r.Use(middleware.Recoverer)

	for _, v := range rs {
		v.Routes(r)
	}

	return r
}

func ListenAndServe(port string) error {
	//Compose the dependicies
	apiRoutes := InitalizeAPIResources(
		NewLinkResource(
			service.NewLinkService(
				service.GetDatabaseConnection(),
				opengraph.NewOGClient(),
			),
		),
	)

	root := chi.NewRouter()
	root.Mount("/v0", apiRoutes)

	log.Infof("Starting Bookmarks API Server on port %s", port)
	return http.ListenAndServe(":"+port, root)
}
