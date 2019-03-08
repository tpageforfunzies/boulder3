// routes.go
package handlers

import (
	u "github.com/tpageforfunzies/boulder/common"
	"github.com/tpageforfunzies/boulder/models"
	"github.com/tpageforfunzies/boulder/services"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


func CreateRoute(c *gin.Context) {

	route := &models.Route{}
	// throw this bitch up in that object
	err := json.NewDecoder(c.Request.Body).Decode(route)
	if err != nil {
		u.Respond(c.Writer, u.Message(false, string(err.Error())))
		return
	}

	resp := services.CreateRoute(route)
	if !resp["status"].(bool) {
		c.JSON(http.StatusForbidden, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
	return
}

func GetRoute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		u.Respond(c.Writer, u.Message(false, "error in your request"))
		return
	}

	route := services.GetRoute(id)
	if route == nil {
		resp := u.Message(false, "could not find route")
		c.JSON(http.StatusNotFound, resp)
		return
	}
	resp := u.Message(true, "success")
	resp["data"] = route
	c.JSON(http.StatusOK, resp)
	return

}

func GetRoutesForUser(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		u.Respond(c.Writer, u.Message(false, "error in your request"))
		return
	}

	data := services.GetRoutesByUserId(id)
	if len(data) == 0 {
		resp := u.Message(false, "could not find routes")
		c.JSON(http.StatusNotFound, resp)
		return
	}
	resp := u.Message(true, "success")
	resp["data"] = data
	c.JSON(http.StatusOK, resp)
	return
}