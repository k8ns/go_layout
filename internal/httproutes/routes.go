package httproutes

import (
	"article/internal/endpoints"
	"encoding/json"
	"net/http"
)

func GetRoutes(addE endpoints.AddArticle) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /{$}", welcome())
	mux.Handle("POST /articles", addArticle(addE))



	return mux
}

func welcome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]string{
			"app": "articles-api",
		})
	}
}

func addArticle(addE endpoints.AddArticle) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := endpoints.AddArticleRequest{}

		w.Header().Add("Content-Type", "applicaton/json")
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		}

		resp := addE.Do(req)

		if resp.Error != "" {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusCreated)
		}

		_ = json.NewEncoder(w).Encode(resp)
	}
}
