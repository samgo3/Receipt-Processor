package main

import (
	"fmt"
	"net/http"
	"receipt-processor/internal/handler"
	"receipt-processor/internal/repository"
	"receipt-processor/internal/router"
	"receipt-processor/internal/service"

	"receipt-processor/internal/utils"

	"github.com/spf13/viper"
)

// @title Fetch Receipt Processor
// @version 0.1
// @host localhost:5555
// @BasePath /
func main() {
	utils.LoadConfig()

	receiptRepo := repository.NewKVRepo()
	receiptService := service.NewReceiptService(receiptRepo)
	receiptHandler := handler.NewReceiptHandler(receiptService)

	router := router.RegisterRoutes(receiptHandler)

	host := viper.GetString("server.host")
	port := viper.GetString("server.port")
	addr := fmt.Sprintf("%s:%s", host, port)

	utils.GetLogger().Info("Starting server on  http://" + addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		panic(err)
	}
}
