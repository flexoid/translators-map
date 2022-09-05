package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

func encodeJSONResponse(i interface{}, status int, w http.ResponseWriter, logger *zap.SugaredLogger) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if status != 0 {
		w.WriteHeader(status)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	err := json.NewEncoder(w).Encode(i)
	if err != nil {
		err = fmt.Errorf("failed to encode JSON response: %w", err)
	}

	return err
}
