package apis

import (
	"net/http"
	"rodolfosrd/models"

	"github.com/gin-gonic/gin"
)

func ItemsIndex(c *gin.Context) {
	s := models.Item{Title: "sea", Notes: "nnn"}
	/*db, _ := c.Get("db")
	conn := db.(pgx.Conn)
	items, err := models.GetAllItems(&conn)
	if err != nil {
		fmt.Println(err)
	}*/
	c.JSON(http.StatusOK, gin.H{
		"message": "Hola" + s.Title,
	})
}
