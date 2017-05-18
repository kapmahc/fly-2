package site

import "github.com/gin-gonic/gin"

// Mount web mount-points
func (p *Plugin) Mount(rt *gin.Engine) {
	rt.GET("/install", p.Wrapper.Handle(p.mustDatabaseEmpty), p.Wrapper.HTML("site/install", p.getInstall))
}
