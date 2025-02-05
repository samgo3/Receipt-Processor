package router

import (
	"net/http"
	"receipt-processor/internal/handler"
	"receipt-processor/internal/middleware"

	_ "receipt-processor/docs"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Route struct {
	Path        string
	Handler     http.Handler
	Methods     []string
	Middlewares []alice.Constructor
}

func routeGroups(receiptHandler *handler.ReceiptHandler) []Route {
	return []Route{
		{
			Path:        "/receipts/process",
			Handler:     http.HandlerFunc(receiptHandler.ProcessReceipt),
			Methods:     []string{http.MethodPost},
			Middlewares: []alice.Constructor{middleware.LoggingMiddleware},
		},
		{
			Path:        "/receipts/{id}/points",
			Handler:     http.HandlerFunc(receiptHandler.GetPointsById),
			Methods:     []string{http.MethodGet},
			Middlewares: []alice.Constructor{middleware.LoggingMiddleware},
		},
	}
}

// RegisterRoutes registers all routes and returns a router
func RegisterRoutes(receiptHandler *handler.ReceiptHandler) http.Handler {

	router := mux.NewRouter().PathPrefix("/").Subrouter()
	for _, route := range routeGroups(receiptHandler) {
		handlerChain := alice.New(route.Middlewares...).Then(route.Handler)
		router.Handle(route.Path, handlerChain).Methods(route.Methods...)
	}
	// Configure Swagger UI
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return router
}
