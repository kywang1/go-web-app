import "github.com/jinzhu/gorm"

func InitDb() *gorm.DB {
	dbConfig := "sslmode=disable host=db port=5432 dbname=todos user=tduser password=tdpass"
	db, err := gorm.Open("postgres", dbConfig)
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	models.RunMigrations(db)

	return db
}