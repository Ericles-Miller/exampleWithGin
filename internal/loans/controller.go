package loans

import "github.com/gin-gonic/gin"


type LoanController struct{
}

func NewLoanController() *LoanController {
	return &LoanController{}
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

}

func (c *LoanController) GetLoan(ctx *gin.Context) {
	ctx.String(200, "FUNCIONOU")
}

func (c *LoanController) GetAllLoans(ctx *gin.Context) {

}

func (c *LoanController) UpdateLoan(ctx *gin.Context) {

}

func (c *LoanController) DeleteLoan(ctx *gin.Context) {

}
