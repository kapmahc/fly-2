package site

import (
	"net/http"

	"github.com/kapmahc/h2o"
)

func (p *Plugin) indexLinks(c *h2o.Context) error {
	var items []Link
	if err := p.Db.Order("loc ASC, sort_order ASC").Find(&items).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, items)

}

type fmLink struct {
	Label     string `form:"label" validate:"required,max=255"`
	Href      string `form:"href" validate:"required,max=255"`
	Loc       string `form:"loc" validate:"required,max=32"`
	SortOrder int    `form:"sortOrder"`
}

func (p *Plugin) createLink(c *h2o.Context) error {
	var fm fmLink
	if err := c.Bind(&fm); err != nil {
		return err
	}
	item := Link{
		Label:     fm.Label,
		Href:      fm.Href,
		Loc:       fm.Loc,
		SortOrder: fm.SortOrder,
	}
	if err := p.Db.Create(&item).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, item)
}

func (p *Plugin) showLink(c *h2o.Context) error {
	var item Link
	if err := p.Db.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, item)
}

func (p *Plugin) updateLink(c *h2o.Context) error {
	var fm fmLink
	if err := c.Bind(&fm); err != nil {
		return err
	}
	if err := p.Db.Model(&Link{}).
		Where("id = ?", c.Param("id")).
		Updates(map[string]interface{}{
			"loc":        fm.Loc,
			"label":      fm.Label,
			"href":       fm.Href,
			"sort_order": fm.SortOrder,
		}).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h2o.H{})
}

func (p *Plugin) destroyLink(c *h2o.Context) error {
	if err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(Link{}).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h2o.H{})
}
