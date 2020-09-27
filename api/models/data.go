package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
  guuid "github.com/google/uuid"
)

type Data struct {
	Id uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Title string `gorm:"size:255;not null;unique" json:"nickname"`
	Uuid4 string `gorm:"size:100;not null;unique" json:"email"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
}

func genUUID() string {
  id := guuid.New()
	return id.String()
}

func (d *Data) Prepare() {
	d.Id = 0
	d.Title = html.EscapeString(strings.TrimSpace(d.Title))
	d.Uuid4 = html.EscapeString(genUUID())
	d.CreatedAt = time.Now()
}

func (d *Data) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if d.Title == "" {
			return errors.New("Required Title")
		}
		return nil

	default:
		if d.Title == "" {
			return errors.New("Required Title")
		}
		return nil
	}
}

func (d *Data) SaveData(db *gorm.DB) (*Data, error) {
	var err error
	err = db.Debug().Create(&d).Error
	if err != nil {
		return &Data{}, err
	}
	return d, nil
}

func (d *Data) FindAllData(db *gorm.DB) (*[]Data, error) {
	var err error
	datas := []Data{}
	err = db.Debug().Model(&Data{}).Limit(100).Find(&datas).Error
	if err != nil {
		return &[]Data{}, err
	}
	return &datas, err
}

func (d *Data) FindDataByID(db *gorm.DB, did uint32) (*Data, error) {
	var err error
	err = db.Debug().Model(Data{}).Where("id = ?", did).Take(&d).Error
	if err != nil {
		return &Data{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Data{}, errors.New("Data Not Found")
	}
	return d, err
}
