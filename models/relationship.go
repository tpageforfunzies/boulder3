// relationship.go
package models

import (
	"github.com/jinzhu/gorm"
)

type Relationship struct {
	gorm.Model
	FollowerID int `json:"location_id"`
	FollowedID int `json:"rating"`
}
