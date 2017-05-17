package auth

import "github.com/gin-gonic/gin"

// Mount web mount-points
func (p *Plugin) Mount(rt *gin.Engine) {
	ung := rt.Group("/users")
	ung.GET("/sign-up", p.Wrapper.HTML("auth/users/sign-up", p.getUsersSignUp))
	ung.GET("/logs", p.Wrapper.HTML("auth/users/logs", p.getUsersSignUp))
}
