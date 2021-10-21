package clients

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sleep.com/cashback/entities"
)

type MySQLManager struct {
	db *gorm.DB
}

func NewMySQLManager() (*MySQLManager, error) {
	dsn := "root:password@tcp(localhost:3306)/sleek?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("Failed to connect with db")
	}

	// Migrate the schema
	db.AutoMigrate(&entities.DealStatus{})

	mySQLManager := &MySQLManager{
		db: db,
	}
	return mySQLManager, nil
}

func (m *MySQLManager) ActivateDeal(id uuid.UUID) {
	var dealStatus entities.DealStatus
	result := m.db.First(&dealStatus, id)
	if result.RowsAffected == 0 {
		dealStatus.ID = id
		dealStatus.UserCount = 1
		m.db.Create(&dealStatus)
	} else {
		dealStatus.UserCount = dealStatus.UserCount + 1
		m.db.Model(&dealStatus).Update("user_count", dealStatus.UserCount)
	}
}

func (m *MySQLManager) GetDealsStatus() []entities.DealStatus {
	var dealsStatus []entities.DealStatus
	_ = m.db.Find(&dealsStatus)
	return dealsStatus
}
