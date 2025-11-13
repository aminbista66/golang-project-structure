package response

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSON(c *gin.Context, status int, data interface{}) {
	c.JSON(status, gin.H{"data": data})
}

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	w.Write(jsonData)
	return nil
}
