package config

import (
	"github.com/davidbenavidez/go-chi-mysql-boilerplate/internal/course"
	"github.com/davidbenavidez/go-chi-mysql-boilerplate/internal/log"
	"github.com/davidbenavidez/go-chi-mysql-boilerplate/internal/student"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

type DBConfiguration struct {
	DBDriver, DBName, Username, Password, Host, Port string
	LogMode                                          bool
}

func init() {
	// Get config from config file
	viper.AddConfigPath("./static")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&serverConfig)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}

func (s *Server) setupDatabase() (*gorm.DB, error) {
	db, err := connectToDatabase(serverConfig.Database)

	db.SingularTable(true)

	// Create Course table
	// db.DropTable(&course.Course{})
	db.AutoMigrate(&course.Course{})

	// Create Student table
	// db.DropTable(&student.Student{})
	db.AutoMigrate(&student.Student{})

	// Create Course Student table
	// db.DropTable(&cs.CourseStudent{})
	// db.AutoMigrate(&cs.CourseStudent{})

	return db, err
}

func connectToDatabase(config DBConfiguration) (*gorm.DB, error) {

	connectionString := config.Username +
		":" +
		config.Password +
		"@tcp(" +
		config.Host +
		":" +
		config.Port +
		")/" +
		config.DBName +
		"?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", connectionString)

	if err != nil {
		log.Fatalf("Error connecting to database: %s", err.Error())
		return nil, err
	}

	log.Infoln("Connected to database")
	return db, nil
}
