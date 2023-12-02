package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var linksEndpoint = "/v1/credentials/links/"
var qrCodeEndpoint = "/qrcode"

func runApi() {
	r := gin.Default()

	api := r.Group("/api/v1")

	api.POST("/create/event", createEvent)
	api.POST("/create/profile", createProfile)

	r.Run(fmt.Sprintf(":%d", config.Port))
}

func createEvent(c *gin.Context) {
	var payload *NewEventRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	eventLinkRequest := EventCredentialLinkRequest{
		SchemaId:            config.EventSchemaId,
		SignatureProof:      true,
		MtProof:             true,
		ClaimLinkExpiration: payload.ValidUntil,
		LimitedClaims:       payload.ClaimLimit,
		CredentialSubject:   payload.EventInfo,
	}

	reqJson, err := json.Marshal(eventLinkRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	createRes, err := sendCreateLinkRequest(reqJson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	qrCodeUrl := config.IssuerApi + linksEndpoint + createRes.Id + qrCodeEndpoint
	png, err := createQrCode(qrCodeUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	c.Header("Content-Type", "image/png")
	c.Header("Content-Length", fmt.Sprint(len(png)))

	c.Data(http.StatusOK, "image/png", png)
}

func createProfile(c *gin.Context) {
	var payload *Profile
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	profileLinkRequest := ProfileCredentialLinkRequest{
		SchemaId:          config.ProfileSchemaId,
		SignatureProof:    true,
		MtProof:           true,
		CredentialSubject: *payload,
	}

	reqJson, err := json.Marshal(profileLinkRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	createRes, err := sendCreateLinkRequest(reqJson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	qrCodeUrl := config.IssuerApi + linksEndpoint + createRes.Id + qrCodeEndpoint
	png, err := createQrCode(qrCodeUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	c.Header("Content-Type", "image/png")
	c.Header("Content-Length", fmt.Sprint(len(png)))

	c.Data(http.StatusOK, "image/png", png)
}

func sendCreateLinkRequest(reqJson []byte) (CreateLinkResponse, error) {
	url := config.IssuerApi + linksEndpoint

	r, err := http.NewRequest("POST", url, bytes.NewBuffer(reqJson))
	if err != nil {
		return CreateLinkResponse{}, err
	}
	r.SetBasicAuth(config.IssuerUname, config.IssuerPass)
	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return CreateLinkResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return CreateLinkResponse{}, fmt.Errorf("request to issuer node failed with status code: %d", res.StatusCode)
	}

	var createRes CreateLinkResponse
	if err = json.NewDecoder(res.Body).Decode(&createRes); err != nil {
		return CreateLinkResponse{}, err
	}
	return createRes, nil
}

func sendGetQrRequest(url string) (string, error) {
	r, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", err
	}
	r.SetBasicAuth(config.IssuerUname, config.IssuerPass)
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("auth QR code request to issuer node failed with status code: %d", res.StatusCode)
	}

	var qrRes QrCodeRequestResponse
	if err = json.NewDecoder(res.Body).Decode(&qrRes); err != nil {
		return "", err
	}

	return qrRes.QrCode, nil
}
