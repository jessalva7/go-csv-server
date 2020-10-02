package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jessalva7/go-csv-file-server/pkg/handlers"
	"github.com/jessalva7/go-csv-file-server/pkg/parsing"
	"log"
	"os"
)

func main() {

	router := gin.Default()
	logger := log.New(os.Stdout, "CSV Server: ", log.LstdFlags)
	parsingService := parsing.NewService(logger)
	uploadHandler := handlers.NewUploadHandler(logger, parsingService)

	router.POST("/uploadCSV", uploadHandler.UploadCSV)
	_ = router.Run(":9074")

}
