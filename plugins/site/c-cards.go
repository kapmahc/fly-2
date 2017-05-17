package site

import (
	"net/http"

	"github.com/kapmahc/h2o"
)

func (p *Plugin) indexCards(c *h2o.Context) error {
	var items []Card
	if err := p.Db.Order("loc ASC, sort_order ASC").Find(&items).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, items)
}

type fmCard struct {
	Loc       string `form:"loc" validate:"required,max=32"`
	Title     string `form:"title" validate:"required,max=255"`
	Summary   string `form:"summary" validate:"required"`
	Href      string `form:"href" validate:"required,max=255"`
	Logo      string `form:"logo" validate:"required,max=255"`
	SortOrder int    `form:"sortOrder"`
}

func (p *Plugin) createCard(c *h2o.Context) error {
	var fm fmCard
	if err := c.Bind(&fm); err != nil {
		return err
	}

	item := Card{
		Title:     fm.Title,
		Logo:      fm.Logo,
		Href:      fm.Href,
		Summary:   fm.Summary,
		SortOrder: fm.SortOrder,
		Loc:       fm.Loc,
		Action:    "buttons.view",
	}
	if err := p.Db.Create(&item).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, item)
}

func (p *Plugin) showCard(c *h2o.Context) error {

	var item Card
	if err := p.Db.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, item)
}

func (p *Plugin) updateCard(c *h2o.Context) error {
	var fm fmCard
	if err := c.Bind(&fm); err != nil {
		return err
	}
	if err := p.Db.Model(&Card{}).
		Where("id = ?", c.Param("id")).
		Updates(map[string]interface{}{
			"href":       fm.Href,
			"title":      fm.Title,
			"logo":       fm.Logo,
			"sort_order": fm.SortOrder,
			"loc":        fm.Loc,
			"summary":    fm.Summary,
		}).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h2o.H{})
}

func (p *Plugin) destroyCard(c *h2o.Context) error {
	if err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(Card{}).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h2o.H{})
}
