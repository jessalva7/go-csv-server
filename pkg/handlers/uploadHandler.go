package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jessalva7/go-csv-file-server/pkg/parsing"
	"log"
	"net/http"
)

type uploadHandler struct {
	log            *log.Logger
	parsingService parsing.Service
}

func (uh uploadHandler) UploadCSV(c *gin.Context) {

	log.Print("Uploading CSV")

	f, err := c.FormFile("csvFile")
	if f.Header.Get("Content-Type") != "text/csv" {

		log.Print("file type is not text/csv")
		c.JSON(http.StatusBadRequest, "file was not a csv file")
		return
	}

	if err != nil {

		log.Print("Couldn't find CSV")
		c.JSON(http.StatusBadRequest, "couldn't find the csvFile")
		return

	}

	file, err := f.Open()
	if err != nil {

		log.Print("Couldn't open CSV")
		c.JSON(http.StatusBadRequest, "couldn't open the csv file")
		return

	}

	defer func() {

		if err != file.Close() {

			log.Print("error closing file")

		}

	}()

	go func() {

		err = uh.parsingService.ParseCSV(file)
		if err != nil {

			log.Print("Couldn't parse CSV")
			c.JSON(http.StatusBadRequest, "couldn't parse the CSV file properly")
			return

		}

	}()


	log.Print("parsed CSV: ", f.Filename)
	c.JSON(http.StatusOK, "Parsed CSV successfully")

}

func NewUploadHandler(logger *log.Logger, parsingService parsing.Service) uploadHandler {

	return uploadHandler{log: logger, parsingService: parsingService}

}
