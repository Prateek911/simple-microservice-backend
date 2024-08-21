package db

import (
	"log"
	"simple-microservice-backend/config"
	"simple-microservice-backend/db/model"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	dbConfig, err := config.NewDBConfig()
	if err != nil {
		log.Fatal("Error Loading database config :", err)
	}
	dsn := config.GetConnectionString(dbConfig)
	log.Println("con string:", dsn)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	pgsDB, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}

	pgsDB.SetMaxIdleConns(int(dbConfig.MaxIdleConnections))
	pgsDB.SetMaxOpenConns(int(dbConfig.MaxOpenConnections))
	pgsDB.SetConnMaxLifetime(time.Duration(dbConfig.MaxLifetime) * time.Second)
}

func Close() error {
	if DB != nil {
		pgsDB, err := DB.DB()
		if err != nil {
			log.Println("Error while closing the database connection:", err)
			return err
		}
		if err := pgsDB.Close(); err != nil {
			log.Println("Error while closing the database connection:", err)
			return err
		}
		log.Println("Database connection closed.")
	}
	return nil
}

func MigrateAndResetDB(db *gorm.DB) {

	db.Migrator().DropTable(&model.AccountMaster{},
		&model.Employee{},
		&model.Owner{},
		&model.Payments{},
		&model.Contact{},
		&model.Contactables{})

	db.AutoMigrate(&model.AccountMaster{},
		&model.Employee{},
		&model.Owner{},
		&model.Payments{},
		&model.Contact{},
		&model.Contactables{})
}
