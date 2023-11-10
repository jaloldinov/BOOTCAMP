package handler

import (
	"app/models"
	"app/pkg/logger"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateTariff godoc
// @Router       /tariff [POST]
// @Summary      CREATES TARIFF
// @Description  CREATES TARIFF BASED ON GIVEN DATA
// @Tags         TARIFF
// @Accept       json
// @Produce      json
// @Param        data  body      models.CreateStaffTarif  true  "tariff data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateStaffTarif(c *gin.Context) {
	var tariff models.CreateStaffTarif
	err := c.ShouldBind(&tariff)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.storage.Tariff().CreateStaffTarif(c.Request.Context(), &tariff)
	if err != nil {
		fmt.Println("error Tariff Create:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Tariff successfully created", "id": resp})
}

// CreateTariff godoc
// @Router       /tariff/{id} [GET]
// @Summary      GET BY ID
// @Description  get tariff by ID
// @Tags         TARIFF
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Tariff ID" format(uuid)
// @Success      200  {object}  models.StaffTarif
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetStaffTarif(c *gin.Context) {
	response := models.StaffTarif{}
	id := c.Param("id")

	ok, err := h.redis.Cache().Get(c.Request.Context(), id, &response)
	if err != nil {
		fmt.Println("not found tariff in redis cache")
	}

	if ok {
		c.JSON(http.StatusOK, response)
		return
	}

	resp, err := h.storage.Tariff().GetStaffTarif(c.Request.Context(), &models.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("error tariff Get:", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": resp})

	err = h.redis.Cache().Create(c.Request.Context(), id, resp, time.Hour)
	if err != nil {
		fmt.Println("error create tariff in redis cache: ", err)
	}
}

// ListTariffs godoc
// @Router       /tariff [GET]
// @Summary      GET  ALL BRANCHS
// @Description  gets all tariff based on limit, page and search by name
// @Tags         TARIFF
// @Accept       json
// @Produce      json
// @Param   limit         query     int        false  "limit"          minimum(1)     default(10)
// @Param   page         query     int        false  "page"          minimum(1)     default(1)
// @Param   search         query     string        false  "search"
// @Success      200  {object}  models.GetAllStaffTarif
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllStaffTarif(c *gin.Context) {
	h.log.Info("request GetAllTariff")
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		h.log.Error("error get page:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid page param")
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		h.log.Error("error get limit:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid page param")
		return
	}

	resp, err := h.storage.Tariff().GetAllStaffTarif(c.Request.Context(), &models.GetAllStaffTarifRequest{
		Page:  page,
		Limit: limit,
		Name:  c.Query("search"),
	})
	if err != nil {
		h.log.Error("error Branch GetAllTariff:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	h.log.Warn("response to GetAllTariff")
	c.JSON(http.StatusOK, resp)
}

// UpdateTariff godoc
// @Router       /tariff/{id} [PUT]
// @Summary      UPDATE TARIFF BY ID
// @Description  UPDATES TARIFF BASED ON GIVEN DATA AND ID
// @Tags         TARIFF
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of tariff" format(uuid)
// @Param        data  body      models.CreateStaffTarif  true  "tariff data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateStaffTarif(c *gin.Context) {
	var tariff models.StaffTarif
	err := c.ShouldBind(&tariff)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	tariff.Id = c.Param("id")
	resp, err := h.storage.Tariff().UpdateStaffTarif(c.Request.Context(), &tariff)
	if err != nil {
		h.log.Error("error Tariff Update:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tariff successfully updated", "id": resp})

	err = h.redis.Cache().Delete(c.Request.Context(), tariff.Id)
	if err != nil {
		fmt.Println("error delete tariff in redis cache: ", err)
	}
}

// DeleteTariff godoc
// @Router       /tariff/{id} [DELETE]
// @Summary      DELETE TARIFF BY ID
// @Description  DELETES TARIFF BASED ON ID
// @Tags         TARIFF
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of tariff" format(uuid)
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteStaffTarif(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.storage.Tariff().DeleteStaffTarif(c.Request.Context(), &models.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error deleting tarif:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "tariff successfully deleted", "id": resp})

	err = h.redis.Cache().Delete(c.Request.Context(), id)
	if err != nil {
		fmt.Println("error delete tariff in redis cache: ", err)
	}
}
