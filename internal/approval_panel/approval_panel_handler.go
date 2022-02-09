package approval_panel

import (
	"fmt"
	"net/http"
	"strconv"

	// "net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type approvalHandler struct {
	approvalService Service
}

func convertToApprovalResponse(p Approval) ApprovalResponse {
	return ApprovalResponse{
		ID:              p.ID,
		Id_Request:      p.Id_Request,
		Id_Employee:     p.Id_Employee,
		Full_Name:       p.Full_Name,
		Leave_Type:      p.Leave_Type,
		Job_Title:       p.Job_Title,
		Division:        p.Division,
		Description:     p.Description,
		Address:         p.Address,
		Start_Date:      p.Start_Date,
		End_Date:        p.End_Date,
		Approval_Status: p.Approval_Status,
	}
}

func NewApprovalHandler(approvalService Service) *approvalHandler {
	return &approvalHandler{approvalService}
}

func (h *approvalHandler) GetAllList(c *gin.Context) {
	approvals, err := h.approvalService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var approvalsResponse []ApprovalResponse

	for _, p := range approvals {
		approvalResponse := convertToApprovalResponse(p)
		approvalsResponse = append(approvalsResponse, approvalResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": approvalsResponse,
	})
}

func (h *approvalHandler) GetListById(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	p, err := h.approvalService.FindByID(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	approvalResponse := convertToApprovalResponse(p)

	c.JSON(http.StatusOK, gin.H{
		"data": approvalResponse,
	})
}

func (h *approvalHandler) Query(c *gin.Context) {
	title := c.Query("title")

	c.JSON(http.StatusOK, gin.H{
		"id":      "1",
		"title":   title,
		"message": "success",
	})
}

func (h *approvalHandler) Create(c *gin.Context) {
	var createApprovalRequest CreateApprovalRequest
	err := c.ShouldBindJSON(&createApprovalRequest)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})

		return
	}

	approval, err := h.approvalService.Create(createApprovalRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": approval,
	})
}

// Go Routine for Form Create List Approval
func (h *approvalHandler) CreateList(c *gin.Context) {
	var createList []CreateApprovalRequest
	err := c.ShouldBindJSON(&createList)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})

		return
	}

	approvals, err := h.approvalService.CreateList(createList)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": approvals,
	})
}

func (h *approvalHandler) Update(c *gin.Context) {
	var updateApprovalRequest UpdateApprovalRequest

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	err := c.ShouldBindJSON(&updateApprovalRequest)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})

		return
	}

	approval, err := h.approvalService.Update(id, updateApprovalRequest)

	approvalResponse := convertToApprovalResponse(approval)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": approvalResponse,
	})
}

func (h *approvalHandler) Delete(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	approval, err := h.approvalService.Delete(id)

	approvalResponse := convertToApprovalResponse(approval)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": approvalResponse,
	})
}

// func (h *approvalHandler) DeleteAllList(c *gin.Context) {
// 	_, err := h.approvalService.DeleteAll()

// 	// approvalResponse := convertToApprovalResponse(approvals)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err,
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"data": "delete all success",
// 	})
// }
