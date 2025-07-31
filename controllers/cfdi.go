package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"facturama-api/models" 
)

var facturamaURL = "https://api.facturama.mx"

// CreateCfdi godoc
// @Summary Crear CFDI
// @Description Genera una nueva factura CFDI (global o normal)
// @Tags CFDI
// @Accept json
// @Produce json
// @Param factura body models.CfdiRequest true "Datos del CFDI (estructura completa)"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/cfdi [post]
func CreateCfdi(c *gin.Context) {
	var factura models.CfdiRequest
	if err := c.ShouldBindJSON(&factura); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos JSON inválidos"})
		return
	}

	jsonData, err := json.Marshal(factura)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al serializar JSON"})
		return
	}

	req, err := http.NewRequest("POST", facturamaURL+"/3/cfdis", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creando solicitud"})
		return
	}
	req.SetBasicAuth(os.Getenv("PRODUCTION_FACTURAMA_USER"), os.Getenv("PRODUCTION_FACTURAMA_PASSWORD"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al contactar Facturama"})
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al interpretar respuesta"})
		return
	}

	if resp.StatusCode >= 400 {
		c.JSON(resp.StatusCode, result)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetCfdis godoc
// @Summary Consultar CFDIs emitidos
// @Description Devuelve una lista de facturas emitidas con filtros
// @Tags CFDI
// @Accept json
// @Produce json
// @Param type query string false "Tipo de CFDI (issued)"
// @Param folioStart query string false "Folio inicial"
// @Param folioEnd query string false "Folio final"
// @Param rfc query string false "RFC"
// @Param taxEntityName query string false "Nombre del receptor"
// @Param dateStart query string false "Fecha inicio (dd/mm/yyyy)"
// @Param dateEnd query string false "Fecha fin (dd/mm/yyyy)"
// @Param status query string false "Estado de factura"
// @Param orderNumber query bool false "Ordenar por folio"
// @Param page query int false "Número de página"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/cfdi [get]
func GetCfdis(c *gin.Context) {
	endpoint := facturamaURL + "/cfdi"

	params := "?type=issued"
	for _, key := range []string{"folioStart", "folioEnd", "rfc", "taxEntityName", "dateStart", "dateEnd", "status", "orderNumber", "page"} {
		if val := c.Query(key); val != "" {
			params += fmt.Sprintf("&%s=%s", key, val)
		}
	}

	req, err := http.NewRequest("GET", endpoint+params, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error construyendo solicitud"})
		return
	}

	req.SetBasicAuth(os.Getenv("PRODUCTION_FACTURAMA_USER"), os.Getenv("PRODUCTION_FACTURAMA_PASSWORD"))
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al contactar Facturama"})
		return
	}
	defer resp.Body.Close()

	var result interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al interpretar respuesta"})
		return
	}

	if resp.StatusCode >= 400 {
		c.JSON(resp.StatusCode, result)
		return
	}

	c.JSON(http.StatusOK, result)
}
