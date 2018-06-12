package api

import (
	"context"
	"net/http"

	"github.com/FourSigma/bookmarks/internal/core"
	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
)

func NewBookmarkResource(ls core.BookmarkService) linkResource {
	return linkResource{
		Codec: DefaultJSONCode,
		link:  ls,
	}
}

type linkResource struct {
	Codec
	link core.BookmarkService
}

func (a linkResource) Routes(r chi.Router) {
	r.Route("/links", func(r chi.Router) {
		r.Get("/", a.List)
		r.Post("/", a.Create)
		//r.Get("/preview", a.Preview)

		//Subroute - /links/:linkId
		r.Route("/{linkId}", func(r chi.Router) {
			r.Use(a.SetBookmarkId)
			r.Get("/", a.Get)
			r.Put("/", a.Update)
			r.Delete("/", a.Delete)
		})
	})

}

func (a linkResource) Create(rw http.ResponseWriter, r *http.Request) {
	l := &core.Bookmark{}
	if err := a.Decode(r.Body, l); err != nil {
		http.Error(rw, "error decoding json", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	//Check if link is valid here
	if err := a.link.Create(r.Context(), l); err != nil {
		http.Error(rw, "error creating link", http.StatusInternalServerError)
		return
	}

}

func (a linkResource) Update(rw http.ResponseWriter, r *http.Request) {
	linkId := r.Context().Value(ContextBookmarkId).(core.BookmarkId)

	l := &core.Bookmark{}
	if err := a.Decode(r.Body, l); err != nil {
		http.Error(rw, "error decoding json", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if err := a.link.Update(r.Context(), linkId, l); err != nil {
		http.Error(rw, "error trying to update link", http.StatusInternalServerError)
		return
	}

}

func (a linkResource) Delete(rw http.ResponseWriter, r *http.Request) {
	linkId := r.Context().Value(ContextBookmarkId).(core.BookmarkId)

	if err := a.link.Delete(r.Context(), linkId); err != nil {
		http.Error(rw, "error trying to delete linkId", http.StatusInternalServerError)
		return
	}

	a.Respond(rw, http.StatusOK, map[string]interface{}{
		"id": linkId,
	})
}

func (a linkResource) Get(rw http.ResponseWriter, r *http.Request) {
	linkId := r.Context().Value(ContextBookmarkId).(core.BookmarkId)

	l := core.Bookmark{}
	var err error
	if l, err = a.link.Get(r.Context(), linkId); err != nil {
		http.Error(rw, "error trying to delete linkId", http.StatusInternalServerError)
		return
	}
	a.Respond(rw, http.StatusOK, l)
}
func (a linkResource) List(rw http.ResponseWriter, r *http.Request) {
	if url := r.FormValue("preview"); url != "" {
		a.Preview(rw, r, url)
		return
	}

	var rs []core.Bookmark
	var err error
	if rs, err = a.link.List(r.Context(), nil); err != nil {
		http.Error(rw, "error trying to list", http.StatusInternalServerError)
		return
	}
	a.Respond(rw, http.StatusOK, rs)

}

func (a linkResource) Preview(rw http.ResponseWriter, r *http.Request, url string) {
	l, err := a.link.GetBookmarkFromURL(r.Context(), url)
	if err != nil {
		http.Error(rw, "error trying to get OG data from url", http.StatusInternalServerError)
		return
	}
	a.Respond(rw, http.StatusOK, l)
}

const ContextBookmarkId = "Context - BookmarkId"

func (a linkResource) SetBookmarkId(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		str := chi.URLParam(r, ContextBookmarkId)
		id, err := uuid.FromString(str)
		if err != nil {
			http.Error(w, "error parsing uuid - linkId url param", http.StatusBadRequest)
			return
		}

		r = r.WithContext(
			context.WithValue(r.Context(), ContextBookmarkId, core.BookmarkId(id)),
		)

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func (a linkResource) Respond(rw http.ResponseWriter, code int, v interface{}) {
	rw.WriteHeader(code)
	switch a.Codec.(type) {
	case jsonCodec:
		rw.Header().Set("Content-Type", "application/json")
	default:
		rw.Header().Set("Content-Type", "text/plain")
	}

	a.Encode(rw, v)
}
