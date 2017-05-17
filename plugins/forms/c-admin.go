package forms

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/kapmahc/fly/web"
	"github.com/kapmahc/h2o"
	"github.com/kapmahc/h2o/i18n"
)

func (p *Plugin) indexForms(c *h2o.Context) error {
	var items []Form
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, items)
}

func (p *Plugin) _parseFields(l, s string) ([]Field, error) {
	var fields []Field
	var items []interface{}
	if err := json.Unmarshal([]byte(s), &items); err != nil {
		return nil, err
	}

	for _, item := range items {
		bad := fmt.Sprintf(
			"%s: %+v",
			p.I18n.T(l, "errors.bad-format"),
			item,
		)
		val, ok := item.(map[string]interface{})
		if !ok {
			return nil, errors.New(bad)
		}
		field := Field{}
		if field.Name, ok = val["name"].(string); !ok {
			return nil, errors.New(bad + " name")
		}
		if field.Label, ok = val["label"].(string); !ok {
			return nil, errors.New(bad + " label")
		}
		if field.Type, ok = val["type"].(string); !ok {
			return nil, errors.New(bad + " type")
		}
		if field.Value, ok = val["value"].(string); !ok {
			return nil, errors.New(bad + " value")
		}

		switch {
		case field.Type == TypeCheckboxs || field.Type == TypeRadios:
			// log.Printf("%v\n", reflect.TypeOf(val["body"]))
			options, ok := val["body"].(map[string]interface{})
			if !ok {
				return nil, errors.New(bad + " body")
			}
			buf, err := json.Marshal(options)
			if err != nil {
				return nil, err
			}
			field.Body = string(buf)
		}

		fields = append(fields, field)
	}
	return fields, nil
}

func (p *Plugin) createForm(c *h2o.Context) error {
	var fm fmForm
	if err := c.Bind(&fm); err != nil {
		return err
	}
	lng := c.Get(i18n.LOCALE).(string)

	deadline, err := time.Parse(web.FormatDateInput, fm.Deadline)
	if err != nil {
		return err
	}
	fields, err := p._parseFields(lng, fm.Fields)
	if err != nil {
		return err
	}
	// log.Printf("FIELDS: %+v", fields)

	item := Form{
		Title:    fm.Title,
		Deadline: deadline,
		Media: web.Media{
			Body: fm.Body,
			Type: fm.Type,
		},
	}
	if err := p.Db.Create(&item).Error; err != nil {
		return err
	}
	for _, field := range fields {
		field.FormID = item.ID
		if err := p.Db.Create(&field).Error; err != nil {
			return err
		}
	}

	return c.JSON(http.StatusOK, item)
}

func (p *Plugin) _buildFields(fm *Form) ([]h2o.H, error) {
	var items []h2o.H
	for _, f := range fm.Fields {
		it := h2o.H{
			"name":  f.Name,
			"label": f.Label,
			"type":  f.Type,
			"value": f.Value,
		}
		switch {
		case f.Type == TypeCheckboxs || f.Type == TypeRadios:
			options := make(map[string]interface{})
			if err := json.Unmarshal([]byte(f.Body), &options); err != nil {
				return nil, err
			}
			it["body"] = options
		}
		items = append(items, it)
	}
	return items, nil
}

func (p *Plugin) _mustSelectForm(c *h2o.Context, l string) error {
	var item Form
	if err := p.Db.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		return err
	}
	if err := p.Db.Model(&item).Association("Fields").Find(&item.Fields).Error; err != nil {
		return err
	}
	if err := p.Db.Model(&item).Association("Records").Find(&item.Records).Error; err != nil {
		return err
	}
	c.Set("item", &item)
	return nil
}

func (p *Plugin) showForm(c *h2o.Context) error {
	item := c.Get("item").(*Form)
	return c.JSON(http.StatusOK, item)
}

type fmForm struct {
	Title    string `form:"title" validate:"required,max=255"`
	Deadline string `form:"deadline" validate:"required"`
	Body     string `form:"body" validate:"required"`
	Type     string `form:"type" validate:"required,max=8"`
	Fields   string `form:"fields" validate:"required"`
}

func (p *Plugin) updateForm(c *h2o.Context) error {
	item := c.Get("item").(*Form)
	var fm fmForm
	if err := c.Bind(&fm); err != nil {
		return err
	}
	lng := c.Get(i18n.LOCALE).(string)

	deadline, err := time.Parse(web.FormatDateInput, fm.Deadline)
	if err != nil {
		return err
	}
	fields, err := p._parseFields(lng, fm.Fields)
	if err != nil {
		return err
	}
	// log.Printf("FIELDS: %+v", fields)

	if err := p.Db.Model(&item).Updates(map[string]interface{}{
		"title":    fm.Title,
		"type":     fm.Type,
		"body":     fm.Body,
		"deadline": deadline,
	}).Error; err != nil {
		return err
	}
	if err := p.Db.Model(item).Association("Fields").Clear().Error; err != nil {
		return err
	}
	for _, f := range fields {
		f.FormID = item.ID
		if err := p.Db.Create(&f).Error; err != nil {
			return err
		}
	}

	return c.JSON(http.StatusOK, h2o.H{})
}

func (p *Plugin) destroyForm(c *h2o.Context) error {
	id := c.Param("id")
	if err := p.Db.Where("form_id = ?", id).Delete(Field{}).Error; err != nil {
		return err
	}
	if err := p.Db.Where("form_id = ?", id).Delete(Record{}).Error; err != nil {
		return err
	}
	if err := p.Db.Where("id = ?", id).Delete(Form{}).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, h2o.H{})
}
