package main

import "github.com/jinzhu/gorm"

type MemberDb struct {
	gorm.Model
	Name string
	PhoneNumber string
	CarNumber string
}

type MemberFile struct {
	Name string
	PhoneNumber string
	CarNumber string
}

