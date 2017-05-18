package site

import "github.com/gin-gonic/gin"

// Mount web mount-points
func (p *Plugin) Mount(rt *gin.Engine) {
	rt.GET("/install", p.Wrapper.Handle(p.mustDatabaseEmpty), p.Wrapper.HTML("site/install", p.getInstall))
	rt.POST("/install", p.Wrapper.Handle(p.mustDatabaseEmpty), p.Wrapper.Form(&fmInstall{}, p.postInstall))

	rt.GET("/leave-words/new", p.Wrapper.HTML("site/leave-words/new", p.newLeaveWord))
	rt.POST("/leave-words", p.Wrapper.Form(&fmLeaveWord{}, p.createLeaveWord))
}
