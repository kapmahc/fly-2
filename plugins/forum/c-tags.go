package forum

import (
	"net/http"

	"github.com/kapmahc/h2o"
)

func (p *Plugin) indexTags(c *h2o.Context) error {
	var tags []Tag
	if err := p.Db.Find(&tags).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, tags)
}

type fmTag struct {
	Name string `form:"name" validate:"required,max=255"`
}

func (p *Plugin) createTag(c *h2o.Context) error {

	var fm fmTag
	if err := c.Bind(&fm); err != nil {
		return err
	}
	t := Tag{Name: fm.Name}
	if err := p.Db.Create(&t).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, t)

}

func (p *Plugin) showTag(c *h2o.Context) error {

	var tag Tag
	if err := p.Db.Where("id = ?", c.Param("id")).First(&tag).Error; err != nil {
		return err
	}
	if err := p.Db.Model(&tag).Association("Articles").Find(&tag.Articles).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, tag)
}

func (p *Plugin) updateTag(c *h2o.Context) error {

	var tag Tag
	if err := p.Db.Where("id = ?", c.Param("id")).First(&tag).Error; err != nil {
		return err
	}

	var fm fmTag
	if err := c.Bind(&fm); err != nil {
		return err
	}

	if err := p.Db.Model(&tag).Update("name", fm.Name).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, h2o.H{})

}

func (p *Plugin) destroyTag(c *h2o.Context) error {
	var tag Tag
	if err := p.Db.Where("id = ?", c.Param("id")).First(&tag).Error; err != nil {
		return err
	}

	if err := p.Db.Model(&tag).Association("Articles").Clear().Error; err != nil {
		return err
	}

	if err := p.Db.Delete(&tag).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h2o.H{})

}
