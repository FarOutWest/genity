package seed

import (
	"log"

  guuid "github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/FarOutWest/genity/api/models"
)

func genUUID() string {
  id := guuid.New()
	return id.String()
}

var datas = []models.Data{
	models.Data{
		Title: "Some Title",
    Uuid4: genUUID(),
	},
  models.Data{
		Title: "Some Title 2",
    Uuid4: genUUID(),
	},
  models.Data{
		Title: "Some Title 3",
    Uuid4: genUUID(),
	},
  models.Data{
		Title: "Some Title 4",
    Uuid4: genUUID(),
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Data{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Data{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range datas {
		err = db.Debug().Model(&models.Data{}).Create(&datas[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
