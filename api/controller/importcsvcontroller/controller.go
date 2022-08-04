package importcsvcontroller

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/TamerB/products-import-service/api/models"
	db "github.com/TamerB/products-import-service/db/sqlc"

	"github.com/TamerB/products-import-service/api/handler"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
)

// ImportRequest provides all functions to import files
type ImportRequest struct {
	File *os.File
}

// ImportCSV receives import requests and starts file imports
func ImportCSV(context *gin.Context) {
	ir := ImportRequest{}
	err := ir.uploadCSV(context.Request, "upload-*.csv")

	if err != nil {
		context.JSON(http.StatusBadRequest, []*models.Error{
			{
				Number: http.StatusBadRequest,
				Text:   err.Error(),
			},
		})
		return
	}

	errs := ir.importRows()

	errorMessages := []*models.Error{}
	for _, err := range errs {
		errorMessages = append(
			errorMessages,
			&models.Error{
				Number: http.StatusBadRequest,
				Text:   err.Error(),
			})
	}

	context.JSON(http.StatusOK, models.BaseResponse{
		Success: true,
		Messages: []*models.Message{
			{Number: http.StatusOK, Type: "Success", Text: "File imported succesfully"},
		},
		Errors: errorMessages,
	})
}

// uploadCSV uploads files
func (ir *ImportRequest) uploadCSV(r *http.Request, destName string) error {

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)

	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		return fmt.Errorf("error Getting the File: %s", err.Error())
	}
	defer file.Close()

	log.Println("info", "Uploaded File: ", handler.Filename)

	//if handler.Filename
	log.Println("info", "File Size: ", handler.Size)
	log.Println("info", "MIME Header: ", handler.Header.Values("Content-Type"))
	csv := false
	for _, val := range handler.Header.Values("Content-Type") {
		if strings.Contains(val, "csv") {
			csv = true
		}
	}
	if !csv {
		return fmt.Errorf("uploaded file is not CSV")
	}

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	ir.File, err = ioutil.TempFile("tmp", destName)

	if err != nil {
		return err
	}

	defer ir.File.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	// write this byte array to our temporary file
	ir.File.Write(fileBytes)

	return nil
}

// importRows reads csv file to an array of rows
func (ir *ImportRequest) importRows() []error {
	in, err := os.Open(ir.File.Name())
	if err != nil {
		return []error{err}
	}
	defer in.Close()

	rows := handler.ImportRows{}

	if err := gocsv.UnmarshalFile(in, &rows.Rows); err != nil {
		return []error{err}
	}

	errs := rows.HandleImport(db.DBStore)

	return errs
}
