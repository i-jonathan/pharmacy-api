package model

import (
	"github.com/i-jonathan/pharmacy-api/config"
	"github.com/speps/go-hashids/v2"
	"strings"
)

func stringInSlice(s string, list []string) bool {
	for _, b := range list {
		if s == b {
			return true
		}
	}
	return false
}

// return true if all strings are valid
func stringValidation(list []string) bool {
	for _, l := range list {
		if l == "" {
			return false
		}
	}
	return true
}

func integerValidation(list []int) bool {
	for _, i := range list {
		if i == 0 {
			return false
		}
	}
	return true
}

//ToHashID converts id to hash ID
func ToHashID(id int) (string, error) {
	hashData := hashids.NewData()

	config2 := config.GetConfig()
	hashData.Salt = config2.HashSalt
	hashData.MinLength = 10

	hashing, err := hashids.NewWithData(hashData)
	if err != nil {
		return "", err
	}

	slug, err := hashing.Encode([]int{id})
	if err != nil {
		return "", err
	}
	return slug, nil
}

//DecodeID converts slug to integer ID
func DecodeID(slug string) (int, error) {
	hashData := hashids.NewData()

	config2 := config.GetConfig()
	hashData.Salt = config2.HashSalt
	hashData.MinLength = 10

	hashing, err := hashids.NewWithData(hashData)
	if err != nil {
		return 0, err
	}
	decoded, err := hashing.DecodeWithError(slug)
	if err != nil {
		return 0, err
	}

	return decoded[0], nil
}

// Valid returns true if permission is valid
func (p *Permission) Valid() bool {
	var models = []string{"account", "order", "return", "category", "product", "supplier"}
	var allowed = []string{"read", "write", "update", "delete"}
	parts := strings.Split(p.Name, ":")

	if stringInSlice(parts[0], models) && stringInSlice(parts[1], allowed) {
		return true
	}
	return false
}

// Valid returns true if role is valid
func (r *Role) Valid() bool {
	return r.Name != ""
}

func (a *Account) Valid() bool {
	toCheck := []string{a.FirstName, a.LastName, a.Password, a.PhoneNumber}

	// Look into doing extra account validation
	// Probably password, email or what not
	return stringValidation(toCheck)
}

func (c *Category) Valid() bool {
	return c.Name != ""
}

func (s *Supplier) Valid() bool {
	toValidate := []string{s.Name, s.Address, s.Email, s.PhoneNumber}
	return stringValidation(toValidate)
}

func (p *Product) Valid() bool {
	stringToValidate := []string{p.Name, p.BarCode}
	intToValidate := []int{p.PurchasePrice, p.SellingPrice, p.ReorderLevel}

	return stringValidation(stringToValidate) && integerValidation(intToValidate)
}

func (p *PaymentMethod) Valid() bool {
	toValidate := []string{p.Name}
	return stringValidation(toValidate)
}

func (i *OrderItem) Valid() bool {
	intToValidate := []int{i.Quantity}
	return integerValidation(intToValidate)
}

func (o *Order) Valid() bool {
	return true
}
