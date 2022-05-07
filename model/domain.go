package model

import "strings"

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
	stringToValudate := []string{p.Name, p.BarCode}
	intToValidate := []int{p.PurchasePrice, p.SellingPrice, p.ReorderLevel}

	return stringValidation(stringToValudate) && integerValidation(intToValidate)
}
