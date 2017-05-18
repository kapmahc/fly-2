package site

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kapmahc/fly/i18n"
)

func (p *Plugin) newLeaveWord(c *gin.Context, v gin.H) error {
	l := c.MustGet(i18n.LOCALE).(string)
	v["title"] = p.I18n.T(l, "site.leave-words.new.title")
	return nil
}

type fmLeaveWord struct {
	Body string `form:"body" binding:"required,max=2048"`
	Type string `form:"type" binding:"required,max=16"`
}

func (p *Plugin) createLeaveWord(c *gin.Context, o interface{}) error {
	l := c.MustGet(i18n.LOCALE).(string)
	fm := o.(*fmLeaveWord)
	item := LeaveWord{
		Body: fm.Body,
		Type: fm.Type,
	}
	if err := p.Db.Create(&item).Error; err != nil {
		return err
	}
	c.JSON(http.StatusOK, gin.H{"message": p.I18n.T(l, "success")})
	return nil
}

// TODO

func (p *Plugin) indexLeaveWords(c *gin.Context) error {
	var items []LeaveWord
	if err := p.Db.Order("created_at DESC").Find(&items).Error; err != nil {
		return err
	}
	c.JSON(http.StatusOK, items)
	return nil
}

func (p *Plugin) destroyLeaveWord(c *gin.Context) error {
	if err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(LeaveWord{}).Error; err != nil {
		return err
	}
	c.JSON(http.StatusOK, gin.H{})
	return nil
}
