package mail

import (
	"net/http"

	"github.com/kapmahc/h2o"
)

func (p *Plugin) indexDomains(c *h2o.Context) error {

	var items []Domain
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, items)
}

type fmDomain struct {
	Name string `form:"name" validate:"required,max=255"`
}

func (p *Plugin) createDomain(c *h2o.Context) error {

	var fm fmDomain
	if err := c.Bind(&fm); err != nil {
		return err
	}

	if err := p.Db.Create(&Domain{
		Name: fm.Name,
	}).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h2o.H{})
}

func (p *Plugin) showDomain(c *h2o.Context) error {
	var item Domain
	if err := p.Db.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, item)
}

func (p *Plugin) updateDomain(c *h2o.Context) error {
	var fm fmDomain
	if err := c.Bind(&fm); err != nil {
		return err
	}

	var item Domain
	if err := p.Db.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		return err
	}

	if err := p.Db.Model(&item).
		Updates(map[string]interface{}{
			"name": fm.Name,
		}).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, h2o.H{})
}

func (p *Plugin) destroyDomain(c *h2o.Context) error {
	if err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(Domain{}).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h2o.H{})
}
