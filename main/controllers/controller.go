package controllers

import (
	"../models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func GENDER(value string) string {
	db, err := gorm.Open("postgres", "user=postgres dbname=gorm password=mglo0704 sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	db.CreateTable(&models.Profile{})
	gender := models.Gender{}
	a := models.Profile{}
	a.FirstName=value
	if gender.Gender == a.Gender {
		db.Where("first_name =?", a.FirstName).Find(&a).Select([]string{"gender"}).Scan(&gender)
		fmt.Printf("\n%v\n",a.Gender)
	}
	return gender.Gender
}
func COUNT(value string) int64 {
	db, err := gorm.Open("postgres", "user=postgres dbname=gorm password=mglo0704 sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	db.CreateTable(&models.Profile{})
	a := models.Profile{}
	a.FirstName=value
	//db.Table("product").Where("status = ?", status).Count(&result)
	var result int64
	db.Table("profiles").Where("first_name=?",a.FirstName).Count(&result)
	return result
}
func PROBABILITY(value string) string {
	db, err := gorm.Open("postgres", "user=postgres dbname=gorm password=mglo0704 sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	db.CreateTable(&models.Profile{})
	a := models.Profile{}
	gender := models.Gender{}
	a.FirstName = value
	if gender.Gender == a.Gender {
		db.Where("first_name =?", a.FirstName).Find(&a).Select([]string{"gender"}).Scan(&gender)
		fmt.Printf("\n%v\n", a.Gender)
	}
	if gender.Gender == "male" {
		var result int64
		var z int64
		db.Table("profiles").Where("first_name=?", a.FirstName).Count(&result)
		db.Table("profiles").Where("first_name=? and gender=?", a.FirstName, "male").Count(&z)
		var e = float32(result)
		var t = float32(z)
		c := t / e
		s := fmt.Sprintf("%0.2f",c)
		return s
	} else {
		var result int64
		var z int64
		db.Table("profiles").Where("first_name=?", a.FirstName).Count(&result)
		db.Table("profiles").Where("first_name=? and gender=?", a.FirstName, "female").Count(&z)
		var e = float32(result)
		var t = float32(z)
		c := t / e
		s := fmt.Sprintf("%0.2f",c)
		return s
	}
}
