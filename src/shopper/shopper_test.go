package shopper

import (
	"testing"
	"math/rand"
)

var s = &Shopper{
Id: int(rand.Int31n(2500)),
Email: "akashshinde159@gmail.com",
Address: "Pune",
AreaName: "Bibawewadi",
BankAccountName: "Hack this Account",
BankAccountNo: "123456789",
BankIFSCCode: "SBINH0013666",
ContactNo: 7276,
Description: "This is description",
Latitude: "15.12564",
Longitude: "34.14654",
PanNo: "2132489",
TinNo: "1154654",
ShopOwner: "Akash Shinde",
UserName: "akashshinde",
VatNo: "1234654",
UserPass: "this_should_be_encrypted",
}

func TestShopper_Add(t *testing.T) {
	err := s.Add()
	if err != nil {
		t.Error("Error should be null")
	}
}

func TestShopper_Remove(t *testing.T) {
	err := s.Remove()
	if err != nil {
		t.Error("Error should be null")
	}
}

func TestShopper_Update(t *testing.T) {

}