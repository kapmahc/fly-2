package site

import (
	"net/http"

	"github.com/kapmahc/fly/web"
	"github.com/kapmahc/h2o"
)

func (p *Plugin) indexPosts(c *h2o.Context) error {
	var items []Post
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, items)
}

func (p *Plugin) showPost(c *h2o.Context) error {
	var item Post
	if err := p.Db.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, item)
}

type fmPost struct {
	Title string `form:"title" validate:"required,max=255"`
	Body  string `form:"body" validate:"required"`
	Type  string `form:"type" validate:"required,max=8"`
}

func (p *Plugin) createPost(c *h2o.Context) error {
	var fm fmPost
	if err := c.Bind(&fm); err != nil {
		return err
	}
	item := Post{
		Media: web.Media{Type: fm.Type, Body: fm.Body},
		Title: fm.Title,
	}
	if err := p.Db.Create(&item).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, item)
}

func (p *Plugin) updatePost(c *h2o.Context) error {
	var fm fmPost
	if err := c.Bind(&fm); err != nil {
		return err
	}
	if err := p.Db.Model(&Post{}).
		Where("id = ?", c.Param("id")).
		Updates(map[string]interface{}{
			"body":  fm.Body,
			"type":  fm.Type,
			"title": fm.Title,
		}).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h2o.H{})
}

func (p *Plugin) destroyPost(c *h2o.Context) error {
	if err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(Post{}).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h2o.H{})
}
