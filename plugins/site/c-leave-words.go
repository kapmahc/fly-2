package site

import (
	"net/http"

	"github.com/kapmahc/h2o"
)

func (p *Plugin) indexLeaveWords(c *h2o.Context) error {
	var items []LeaveWord
	if err := p.Db.Order("created_at DESC").Find(&items).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, items)
}

type fmLeaveWord struct {
	Body string `form:"body" validate:"required,max=2048"`
	Type string `form:"type" validate:"required,max=16"`
}

func (p *Plugin) createLeaveWord(c *h2o.Context) error {
	var fm fmLeaveWord
	if err := c.Bind(&fm); err != nil {
		return err
	}

	item := LeaveWord{
		Body: fm.Body,
		Type: fm.Type,
	}
	if err := p.Db.Create(&item).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h2o.H{})
}

func (p *Plugin) destroyLeaveWord(c *h2o.Context) error {
	if err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(LeaveWord{}).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h2o.H{})
}
