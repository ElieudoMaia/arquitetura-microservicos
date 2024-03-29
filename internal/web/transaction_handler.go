package web

import (
	"encoding/json"
	"net/http"

	"github.com/elieudomaia/ms-wallet-app/internal/usecase/create_transaction"
)

type WebTransactionHandler struct {
	CreateTrasactionUseCase create_transaction.CreateTransactionUseCase
}

func NewTransactionHandler(createTransactionUseCase create_transaction.CreateTransactionUseCase) *WebTransactionHandler {
	return &WebTransactionHandler{
		CreateTrasactionUseCase: createTransactionUseCase,
	}
}

func (h *WebTransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var dto create_transaction.CreateTransactionInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := h.CreateTrasactionUseCase.Execute(r.Context(), dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// json response
		w.Header().Add("Content-Type", "application/json")
		// put response body as { error: "error message" }
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
