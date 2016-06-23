package shopper

import "db_helper"

type Shopper struct{
	Id int                `json:"id"`

// Personal Information
	Name string        `json:"name"`
	ShopOwner string `json:"owner_name"`
	Description string        `json:"description"`
	Address string        `json:"address"`
	ContactNo int        `json:"contact_no"`
	Latitude string        `json:"latitude"`
	Longitude string        `json:"longitude"`
	AreaName string        `json:"area_name"`

// Government Document
	VatNo string        `json:"vat_no"`
	PanNo string        `json:"pan_no"`
	TinNo string        `json:"tin_no"`
	BankAccountName string `json:"bank_account_name"`
	BankAccountNo string        `json:"bank_account_no"`
	BankIFSCCode string        `json:"bank_ifsc_code"`

// Credentials
	Email string        `json:"email"`
	UserName string        `json:"username"`
	UserPass string        `json:"password"`
}

// Creates shopper entry in database
func(s *Shopper) Add() error {
	d := db_helper.SharedConnection()
	return d.Create(s).Error
}

func (s *Shopper ) Remove() error {
	d := db_helper.SharedConnection()
	return d.Delete(s).Error
}

func (s *Shopper ) Update() error {
	d := db_helper.SharedConnection()
	return d.Update(s).Error
}

// Hack to get this module working first time
func main() {

}