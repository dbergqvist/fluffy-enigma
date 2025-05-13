package main

import (
	"encoding/base64"
	"log"
	"net/http"

	"fluffy-enigma/api"
	"fluffy-enigma/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	if err := config.LoadConfig("config.json"); err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	r := gin.Default()

	// Serve static files
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	// Routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Certificate Signing App",
		})
	})

	// API routes
	api := r.Group("/api")
	{
		api.POST("/submit-certificates", handleSubmitCertificates)
		api.POST("/teacher-sign", handleTeacherSign)
	}

	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func handleSubmitCertificates(c *gin.Context) {
	var request struct {
		Emails []string `json:"emails"`
		JSON   string   `json:"json"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Scrive client
	scriveClient := api.NewScriveClient()

	// Add teacher's email to signatories
	signatories := append(request.Emails, config.AppConfig.TeacherEmail)

	// Create document in Scrive
	response, err := scriveClient.CreateDocument(
		"Course Completion Certificate",
		signatories,
		base64.StdEncoding.EncodeToString([]byte(request.JSON)),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Certificates submitted successfully",
		"documentId": response.DocumentID,
	})
}

func handleTeacherSign(c *gin.Context) {
	var request struct {
		DocumentID string `json:"documentId"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Scrive client
	scriveClient := api.NewScriveClient()

	// Initiate signing process
	if err := scriveClient.InitiateSigning(request.DocumentID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Teacher signing initiated successfully",
	})
}
