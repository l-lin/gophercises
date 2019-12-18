package web

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/l-lin/gophercises/transform/primitive"
)

const (
	nbShapes      = 10
	maxGoroutines = 2
)

var uploadPath = os.TempDir()

// UploadHandler is the handler to upload images
func UploadHandler(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		displayError(c, http.StatusBadRequest, err.Error())
		return
	}

	inFilePath := fmt.Sprintf("%s/%s", uploadPath, fileHeader.Filename)
	if err := c.SaveUploadedFile(fileHeader, inFilePath); err != nil {
		displayError(c, http.StatusInternalServerError, err.Error())
		return
	}

	go performTransformations(inFilePath)

	c.Redirect(http.StatusCreated, "/images")
}

// ImageHandler to handle and display images
func ImageHandler(c *gin.Context) {
	imageName := c.Params.ByName("imageName")
	inFilePath := fmt.Sprintf("%s/%s", uploadPath, imageName)
	if !exists(inFilePath) {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Image '%s' not found", imageName)})
		return
	}
	ext := filepath.Ext(imageName)
	if ext == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image type not recognized"})
		return
	}
	b, err := ioutil.ReadFile(inFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.Data(200, fmt.Sprintf("image/%s", ext[1:]), b)
}

func displayError(c *gin.Context, statusCode int, message string) {
	log.Printf("[ERROR] %s", message)
	c.JSON(statusCode, gin.H{"error": message})
}

func performTransformations(inFilePath string) {
	guard := make(chan struct{}, maxGoroutines)
	for _, m := range primitive.Modes {
		guard <- struct{}{} // would block if guard channel is already filled
		go func(inFilePath string, m primitive.Mode) {
			log.Printf("[INFO] executing primitive in mode '%s'...", m.String())
			if err := primitive.Transform(inFilePath, uploadPath, m, nbShapes); err != nil {
				log.Printf("[ERRROR] %s", err)
			}
			<-guard
		}(inFilePath, m)
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	log.Println(err)
	return true
}
