package api

import (
	"context"
	"net/http"

	"github.com/FourSigma/bookmarks/internal/core"
	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
)

func NewBookmarkResource(ls core.BookmarkService) bookmarkResource {
	return bookmarkResource{
		Codec:    DefaultJSONCode,
		bookmark: ls,
	}
}

type bookmarkResource struct {
	Codec
	bookmark core.BookmarkService
}

func (a bookmarkResource) Routes(r chi.Router) {
	r.Route("/bookmarks", func(r chi.Router) {
		r.Get("/", a.List)
		r.Post("/", a.Create)
		//r.Get("/preview", a.Preview)

		//Subroute - /bookmarks/:bookmarkId
		r.Route("/{bookmarkId}", func(r chi.Router) {
			r.Use(a.SetBookmarkId)
			r.Get("/", a.Get)
			r.Put("/", a.Update)
			r.Delete("/", a.Delete)
		})
	})

}

func (a bookmarkResource) Create(rw http.ResponseWriter, r *http.Request) {
	l := &core.Bookmark{}
	if err := a.Decode(r.Body, l); err != nil {
		http.Error(rw, "error decoding json", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	//Check if bookmark is valid here
	if err := a.bookmark.Create(r.Context(), l); err != nil {
		http.Error(rw, "error creating bookmark", http.StatusInternalServerError)
		return
	}

}

func (a bookmarkResource) Update(rw http.ResponseWriter, r *http.Request) {
	bookmarkId := r.Context().Value(ContextBookmarkId).(core.BookmarkId)

	l := &core.Bookmark{}
	if err := a.Decode(r.Body, l); err != nil {
		http.Error(rw, "error decoding json", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if err := a.bookmark.Update(r.Context(), bookmarkId, l); err != nil {
		http.Error(rw, "error trying to update bookmark", http.StatusInternalServerError)
		return
	}

}

func (a bookmarkResource) Delete(rw http.ResponseWriter, r *http.Request) {
	bookmarkId := r.Context().Value(ContextBookmarkId).(core.BookmarkId)

	if err := a.bookmark.Delete(r.Context(), bookmarkId); err != nil {
		http.Error(rw, "error trying to delete bookmarkId", http.StatusInternalServerError)
		return
	}

	a.Respond(rw, http.StatusOK, map[string]interface{}{
		"id": bookmarkId,
	})
}

func (a bookmarkResource) Get(rw http.ResponseWriter, r *http.Request) {
	bookmarkId := r.Context().Value(ContextBookmarkId).(core.BookmarkId)

	l := core.Bookmark{}
	var err error
	if l, err = a.bookmark.Get(r.Context(), bookmarkId); err != nil {
		http.Error(rw, "error trying to delete bookmarkId", http.StatusInternalServerError)
		return
	}
	a.Respond(rw, http.StatusOK, l)
}
func (a bookmarkResource) List(rw http.ResponseWriter, r *http.Request) {
	if url := r.FormValue("preview"); url != "" {
		a.Preview(rw, r, url)
		return
	}

	var rs []core.Bookmark
	var err error
	if rs, err = a.bookmark.List(r.Context(), nil); err != nil {
		http.Error(rw, "error trying to list", http.StatusInternalServerError)
		return
	}
	a.Respond(rw, http.StatusOK, rs)

}

func (a bookmarkResource) Preview(rw http.ResponseWriter, r *http.Request, url string) {
	l, err := a.bookmark.GetBookmarkFromURL(r.Context(), url)
	if err != nil {
		http.Error(rw, "error trying to get OG data from url", http.StatusInternalServerError)
		return
	}
	a.Respond(rw, http.StatusOK, l)
}

const ContextBookmarkId = "Context - BookmarkId"

func (a bookmarkResource) SetBookmarkId(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		str := chi.URLParam(r, ContextBookmarkId)
		id, err := uuid.FromString(str)
		if err != nil {
			http.Error(w, "error parsing uuid - bookmarkId url param", http.StatusBadRequest)
			return
		}

		r = r.WithContext(
			context.WithValue(r.Context(), ContextBookmarkId, core.BookmarkId(id)),
		)

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func (a bookmarkResource) Respond(rw http.ResponseWriter, code int, v interface{}) {
	rw.WriteHeader(code)
	switch a.Codec.(type) {
	case jsonCodec:
		rw.Header().Set("Content-Type", "application/json")
	default:
		rw.Header().Set("Content-Type", "text/plain")
	}

	a.Encode(rw, v)
}
