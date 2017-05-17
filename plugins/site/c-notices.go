package site

import (
	"net/http"

	"github.com/kapmahc/fly/web"
	"github.com/kapmahc/h2o"
)

func (p *Plugin) indexNotices(c *h2o.Context) error {
	var items []Notice
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, items)
}

type fmNotice struct {
	Body string `form:"body" validate:"required"`
	Type string `form:"type" validate:"required,max=8"`
}

func (p *Plugin) createNotice(c *h2o.Context) error {
	var fm fmNotice
	if err := c.Bind(&fm); err != nil {
		return err
	}
	item := Notice{
		Media: web.Media{Type: fm.Type, Body: fm.Body},
	}
	if err := p.Db.Create(&item).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, item)
}

func (p *Plugin) showNotice(c *h2o.Context) error {
	var item Notice
	if err := p.Db.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, item)
}

func (p *Plugin) updateNotice(c *h2o.Context) error {
	var fm fmNotice
	if err := c.Bind(&fm); err != nil {
		return err
	}
	if err := p.Db.Model(&Notice{}).
		Where("id = ?", c.Param("id")).
		Updates(map[string]interface{}{
			"body": fm.Body,
			"type": fm.Type,
		}).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h2o.H{})
}

func (p *Plugin) destroyNotice(c *h2o.Context) error {
	if err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(Notice{}).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h2o.H{})
}
