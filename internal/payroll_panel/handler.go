package payroll_panel

import (
	"fmt"
	"net/http"
	"strconv"

	// "net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type payrollHandler struct {
	payrollService Service
}

func convertToPayrollResponse(p Payroll) PayrollResponse {
	return PayrollResponse{
		ID:             p.ID,
		Id_Payment:     p.Id_Payment,
		Id_Employee:    p.Id_Employee,
		Full_Name:      p.Full_Name,
		Job_Title:      p.Job_Title,
		Payment_Period: p.Payment_Period,
		Payment_Date:   p.Payment_Date,
		Payment_Status: p.Payment_Status,
		Basic_Salary:   p.Basic_Salary,
		Bpjs:           p.Bpjs,
		Tax:            p.Tax,
		Total_Salary:   p.Total_Salary,
	}
}

func NewPayrollHandler(payrollService Service) *payrollHandler {
	return &payrollHandler{payrollService}
}

func (h *payrollHandler) GetAllList(c *gin.Context) {
	payrolls, err := h.payrollService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var payrollsResponse []PayrollResponse

	for _, p := range payrolls {
		payrollResponse := convertToPayrollResponse(p)
		payrollsResponse = append(payrollsResponse, payrollResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": payrollsResponse,
	})
}

func (h *payrollHandler) GetListById(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	p, err := h.payrollService.FindByID(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	payrollResponse := convertToPayrollResponse(p)

	c.JSON(http.StatusOK, gin.H{
		"data": payrollResponse,
	})
}

func (h *payrollHandler) QueryList(c *gin.Context) {
	title := c.Query("title")

	c.JSON(http.StatusOK, gin.H{
		"id":      "1",
		"title":   title,
		"message": "success",
	})
}

func (h *payrollHandler) CreateList(c *gin.Context) {
	var createPayrollRequest CreatePayrollRequest
	err := c.ShouldBindJSON(&createPayrollRequest)
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

	payroll, err := h.payrollService.Create(createPayrollRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": payroll,
	})
}

func (h *payrollHandler) UpdateList(c *gin.Context) {
	var updatePayrollRequest UpdatePayrollRequest

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	err := c.ShouldBindJSON(&updatePayrollRequest)
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

	payroll, err := h.payrollService.Update(id, updatePayrollRequest)

	payrollResponse := convertToPayrollResponse(payroll)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": payrollResponse,
	})
}

func (h *payrollHandler) DeleteList(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	payroll, err := h.payrollService.Delete(id)

	payrollResponse := convertToPayrollResponse(payroll)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": payrollResponse,
	})
}

// func (h *payrollHandler) DeleteAllList(c *gin.Context) {
// 	_, err := h.payrollService.DeleteAll()

// 	// payrollResponse := convertToPayrollResponse(payrolls)

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
