package router

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func FileUpload(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, "Not allowed")
		return
	}
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	successmsg := UploadFiles(file, header)

	c.JSON(http.StatusOK, gin.H{
		"message": successmsg,
	})
	//internal.DocsToPdf()
}

func UploadFiles(file multipart.File, header *multipart.FileHeader) string {
	log.Println("uploading a file.....")

	//Getting the input file from the request
	defer file.Close()

	//Getting the filename
	filename := filepath.Base(header.Filename)
	//Creating a file and uploading the file inside the uploads folder
	out, err := os.Create(filepath.Join("uploads", filename))
	if err != nil {
		log.Println("serror from creating the file", err)
	}
	defer out.Close()

	// Copy contents from uploaded file to new file on disk
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	// Convert the PDF to DOCX
	successmsg := "File uploaded successfully"
	// Return success response
	log.Println(successmsg)
	return successmsg
}
