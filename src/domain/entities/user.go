package entities

import (
	"encoding/hex"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/sha3"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey; index; unique; type:uuid;"`
	Name      string    `gorm:"not null; default:null"`
	Email     string    `gorm:"not null; default:null; unique"`
	Password  string    `gorm:"not null; default:null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (User) TableName() string { return "users" }

func encryptPassword(ps string) (pwd string) {
	hash := sha3.Sum256([]byte(ps))
	pwd = hex.EncodeToString(hash[:])
	return pwd
}
