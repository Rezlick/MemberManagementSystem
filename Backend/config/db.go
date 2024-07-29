package config

import (
   "fmt"
   "time"
   "github.com/rezlick/MemberManagementSystem/entity"
   "gorm.io/driver/sqlite"
   "gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
   return db
}

func ConnectionDB() {
   database, err := gorm.Open(sqlite.Open("mms.db?cache=shared"), &gorm.Config{})
   if err != nil {
       panic("failed to connect database")
   }
   fmt.Println("connected database")
   db = database
}

func SetupDatabase() {
   db.AutoMigrate(
       &entity.Users{},
       &entity.Roles{},
   )
   RoleAdmin := entity.Roles{Role: "Admin"}
   RoleMember := entity.Roles{Role: "Member"}
   db.FirstOrCreate(&RoleAdmin, &entity.Roles{Role: "Admin"})
   db.FirstOrCreate(&RoleMember, &entity.Roles{Role: "Member"})

   hashedPassword, _ := HashPassword("1579")

   BirthDay, _ := time.Parse("2006-01-02", "1988-11-12")

   User := &entity.Users{
       FirstName: "Test",
       LastName:  "User",
       Email:     "testuser@gmail.com",
       Age:       99,
	   PhoneNumber: "0987654321",
       Password:  hashedPassword,
       BirthDay:  BirthDay,
       RoleID:  1,
   }
   db.FirstOrCreate(User, &entity.Users{
       Email: "testuser@gmail.com",
   })
}