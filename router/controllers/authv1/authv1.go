package authv1

import (
	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/router/controllers/authv1/problem"
	"github.com/open-fightcoder/oj-web/router/controllers/authv1/submit"
)

func Register(router *gin.RouterGroup) {
	problemRouter := router.Group("/problem")
	problem.RegisterProblem(problemRouter)
	problem.RegisterCode(problemRouter)
	problem.RegisterCollection(problemRouter)
	problem.RegisterUserProblem(problemRouter)

	RegisterAccount(router)
	rankRouter := router.Group("/rank")
	RegisterRank(rankRouter)

	submitRouter := router.Group("/submit")
	submit.RegisterSubmit(submitRouter)
}
