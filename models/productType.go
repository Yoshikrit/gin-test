package models

import (

)

type ProductTypeEntity struct {
    Id int      `gorm:"primaryKey; column:prodtype_code;"`
    Name string `gorm:"not null;   column:prodtype_name;"`
}

//make it know table name from database instead of gorm convention
func (p ProductTypeEntity) TableName() string {
	return "producttype"
}

type ProductType struct {
    Id     		 int    	`json:"ProdType_Id"`
    Name         string 	`json:"ProdType_Name"`
}

type ProductTypeCreate struct {
    Id     		 int    	`json:"ProdType_Id"      binding:"required,gte=0"`
    Name         string 	`json:"ProdType_Name"    binding:"required,max=40"`
}

type ProductTypeUpdate struct {
    Name         string 	`json:"ProdType_Name"    binding:"required,max=40"`
}

