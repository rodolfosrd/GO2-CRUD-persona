package apis

import (
	"net/http"
	"rodolfosrd/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PersonsIndex(c *gin.Context) {
	/*s := models.Person{Name: "sea",Age: "nnn"}
	c.JSON(http.StatusOK, gin.H{
		"message": "Hola"+s.Name,
	})
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)
	items, err := models.GetAllItems(&conn)
	if err != nil {
		fmt.Println(err)
	}*/
	var lis []models.Person
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	conn.Find(&lis)
	c.JSON(http.StatusOK, lis)
}
