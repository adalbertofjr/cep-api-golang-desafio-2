package entity

import (
	"context"
	"io"
	"net/http"
)

func FetchData(url string, w http.ResponseWriter, r *http.Request) string {
	ctx := context.Background()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		http.Error(w, "404 not found", http.StatusNotFound)
		return "404 not found"

	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "404 not found", http.StatusNotFound)
		return "404 not found"
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "404 not found", http.StatusNoContent)
		return "Content not found"
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "Internal Server Error"
	}

	w.Header().Set("Content-Type", "application/json") // Define o tipo de conteúdo como JSON
	w.WriteHeader(http.StatusOK)                       // Retorna o status 200 OK
	// json.NewEncoder(w).Encode(body)
	// fmt.Fprint(w, string(body))
	return string(body)
}
