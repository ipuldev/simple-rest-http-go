package migration

import(
	"github.com/sobatfillah/model"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB){
	db.AutoMigrate(&model.User{})
}