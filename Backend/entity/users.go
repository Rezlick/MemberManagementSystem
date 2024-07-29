package entity

import (
   "time"
   "gorm.io/gorm"
)

type Users struct {
   gorm.Model
   FirstName string    `json:"first_name"`
   LastName  string    `json:"last_name"`
   Email     string    `json:"email"`
   PhoneNumber string  `json:"phone_number"`
   Age       uint8     `json:"age"`
   Password  string    `json:"-"`
   BirthDay  time.Time `json:"birthday"`
   RoleID    uint      `json:"role_id"`
   Role    *Roles      `gorm:"foreignKey: role_id" json:"role"`
}