package entities

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID `gorm:"->; type: UUID DEFAULT gen_random_uuid(); not null; unique; primaryKey"`
	Name      string    `gorm:"<-; type: VARCHAR (50); not null"`
	Email     string    `gorm:"<-; type: VARCHAR (50); not null; unique"`
	Birthdate time.Time
	Role      string
	CreatedAt time.Time `gorm:"<-create; index"`
	UpdatedAt time.Time
}
