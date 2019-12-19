package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/l-lin/gophercises/transform/web"
	"github.com/spf13/cobra"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Launch web server to upload and transform an image",
		Run:   runServe,
	}
	port int
)

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntVarP(&port, "port", "p", 3000, "the port to start the web server on")
}

func runServe(cmd *cobra.Command, args []string) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})
	r.GET("/images/:imageName", web.ImagesHandler)
	r.POST("/io/images", web.UploadHandler)
	r.GET("/io/images/:imageName", web.ImageHandler)

	// Start the server
	log.Printf("Server started on port %d\n", port)
	log.Fatal(r.Run(fmt.Sprintf(":%d", port)))
}
