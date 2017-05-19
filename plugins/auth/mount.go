package auth

import "github.com/gin-gonic/gin"

// Mount web mount-points
func (p *Plugin) Mount(rt *gin.Engine) {
	ung := rt.Group("/users")
	ung.GET("/sign-up", p.Wrapper.HTML("auth/users/sign-up", p.getUsersSignUp))
	ung.POST("/sign-up", p.Wrapper.Form(&fmSignUp{}, p.postUsersSignUp))
	ung.GET("/sign-in", p.Wrapper.HTML("auth/users/sign-in", p.getUsersSignIn))
	ung.POST("/sign-in", p.Wrapper.Form(&fmSignIn{}, p.postUsersSignIn))
	ung.GET("/confirm", p.Wrapper.HTML("auth/users/email-form", p.getUsersEmailForm("confirm")))
	ung.POST("/confirm", p.Wrapper.Form(&fmEmail{}, p.postUsersConfirm))
	ung.GET("/confirm/:token", p.Wrapper.Handle(p.getUsersConfirm))
	ung.GET("/unlock", p.Wrapper.HTML("auth/users/email-form", p.getUsersEmailForm("unlock")))
	ung.POST("/unlock", p.Wrapper.Form(&fmEmail{}, p.postUsersUnlock))
	ung.GET("/unlock/:token", p.Wrapper.Handle(p.getUsersUnlock))
	ung.GET("/forgot-password", p.Wrapper.HTML("auth/users/email-form", p.getUsersEmailForm("forgot-password")))
	ung.POST("/forgot-password", p.Wrapper.Form(&fmEmail{}, p.postUsersForgotPassword))
	ung.GET("/reset-password/:token", p.Wrapper.HTML("auth/users/reset-password", p.getUsersResetPassword))
	ung.POST("/reset-password", p.Wrapper.Form(&fmResetPassword{}, p.postUsersResetPassword))

	umg := rt.Group("/users", p.Jwt.MustSignInMiddleware)
	umg.GET("/self", p.Wrapper.HTML("auth/users/self", p.getUsersSelf))
	umg.POST("/info", p.Wrapper.Form(&fmInfo{}, p.postUsersInfo))
	umg.POST("/change-password", p.Wrapper.Form(&fmChangePassword{}, p.postUsersChangePassword))
	umg.DELETE("/sign-out", p.Wrapper.Handle(p.deleteUsersSignOut))

	rt.GET("/users", p.Wrapper.HTML("auth/users/index", p.indexUsers))
}
