package model

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/i-jonathan/pharmacy-api/config"
	appError "github.com/i-jonathan/pharmacy-api/error"
	"github.com/speps/go-hashids/v2"
	"golang.org/x/crypto/argon2"
	"regexp"
	"strings"
	"time"
)

type passwordConfig struct {
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}

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
	if id < 1 {
		return "", nil
	}
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

	re := regexp.MustCompile("^.+@.+\\..+$")
	validity := re.MatchString(a.Email)

	if !validity {
		return false
	}

	// Look into doing extra account validation
	// Probably password, email or what not
	return stringValidation(toCheck)
}

func (a *Account) HashPassword() error {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return err
	}

	passConfig := &passwordConfig{
		time:    1,
		memory:  64 * 1024,
		threads: 4,
		keyLen:  32,
	}
	hash := argon2.IDKey([]byte(a.Password), salt, passConfig.time, passConfig.memory, passConfig.threads, passConfig.keyLen)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Format for storing argon2id in database: argon2 version, memory, time,
	// number of threads, salt and hash encoded in base 64
	format := "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
	full := fmt.Sprintf(format, argon2.Version, passConfig.memory, passConfig.time, passConfig.threads, b64Salt, b64Hash)
	a.Password = full
	return nil
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

func (a *Account) CreateToken() (string, error) {
	hash, err := ToHashID(a.ID)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"hash":           hash,
		"StandardClaims": jwt.StandardClaims{ExpiresAt: time.Now().Add(16 * time.Hour).Unix()},
	})
	config2 := config.GetConfig()
	hmacSecret := []byte(config2.HMAC)
	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (a *Auth) ComparePassword(hash string) (bool, error) {
	parts := strings.Split(hash, "$")
	passConfig := &passwordConfig{}

	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &passConfig.memory, &passConfig.time, &passConfig.threads)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	passConfig.keyLen = uint32(len(decodedHash))

	comparisonHash := argon2.IDKey([]byte(a.Password), salt, passConfig.time, passConfig.memory, passConfig.threads, passConfig.keyLen)

	return subtle.ConstantTimeCompare(decodedHash, comparisonHash) == 1, nil
}

func ParseToken(tokenString string) (map[string]interface{}, error) {
	// TODO check if token is blacklisted
	config2 := config.GetConfig()
	hmacSecret := []byte(config2.HMAC)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, appError.Unauthorized
		}
		return hmacSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
