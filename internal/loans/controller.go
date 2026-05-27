package loans

import (
	"exampleWithGin/internal/loans/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


type LoanController struct{
	loansService models.LoansService
}

func NewLoanController(loansService models.LoansService) *LoanController {
	return &LoanController{loansService: loansService}
}


func (c *LoanController) RegisterRoutes(r *gin.Engine) {
	loans := r.Group("/loans")
	{
		loans.GET("/", c.GetAllLoans)
		loans.GET("/:id", c.GetLoan)
		loans.POST("", c.CreateLoan)
		loans.PUT("/:id", c.UpdateLoan)
		loans.DELETE("/:id", c.DeleteLoan)
	}
}

func (c *LoanController) CreateLoan(ctx *gin.Context) {
	var loan models.Loan

	if err := ctx.ShouldBindJSON(&loan); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
	}

	createdLoan, err := c.loansService.CreateLoan(&loan)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create loan"})
		return
	}

	ctx.JSON(http.StatusCreated, createdLoan)
}

func (c *LoanController) GetLoan(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	loan, err := c.loansService.GetLoan(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}

	ctx.JSON(http.StatusOK, loan)
}

func (c *LoanController) GetAllLoans(ctx *gin.Context) {
	loans, err := c.loansService.GetAllLoans()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve loans"})
		return
	}

	ctx.JSON(http.StatusOK, loans)
}

func (c *LoanController) UpdateLoan(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	var loan models.Loan

	if err := ctx.ShouldBindJSON(&loan); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
	}

	updatedLoan, err := c.loansService.UpdateLoan(id, &loan)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update loan"})
		return
	}

	ctx.JSON(http.StatusOK, updatedLoan)

}

func (c *LoanController) DeleteLoan(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	err = c.loansService.DeleteLoan(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}

	ctx.Status(http.StatusNoContent)
}
