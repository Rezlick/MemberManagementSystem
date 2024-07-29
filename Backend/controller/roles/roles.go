package roles

import (
   "net/http"
   "github.com/rezlick/MemberManagementSystem/config"
   "github.com/rezlick/MemberManagementSystem/entity"
   "github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
   db := config.DB()
   var roles []entity.Roles
   db.Find(&roles)
   c.JSON(http.StatusOK, &roles)
}