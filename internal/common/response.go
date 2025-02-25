package common

import (
	"encoding/json"
	"net/http"
)

// ResponseJSON принимает http.ResponseWriter и объект данных, устанавливает заголовок
// и записывает JSON-представление данных в ответ.
func ResponseJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
	}
}
