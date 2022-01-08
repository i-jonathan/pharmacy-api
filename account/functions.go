package account

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/argon2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	name := os.Getenv("database_name")
	pass := os.Getenv("database_pass")
	user := os.Getenv("database_user")
	host := os.Getenv("database_host")
	ssl := os.Getenv("databse_ssl")
	port := os.Getenv("databse_port")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, user, name, pass, ssl)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("err")
		return nil
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Println(err)
	}

	return db
}

func generatePasswordHash(password string) (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	config := &passwordConfig{
		time:    1,
		memory:  64 * 1024,
		threads: 4,
		keyLen:  32,
	}

	hash := argon2.IDKey([]byte(password), salt, config.time, config.memory, config.threads, config.keyLen)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	format := "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
	full := fmt.Sprintf(format, argon2.Version, config.memory, config.time, config.threads, b64Salt, b64Hash)

	return full, nil
}

func comparePassword(password, hash string) (bool, error) {
	parts := strings.Split(hash, "$")
	config := &passwordConfig{}

	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &config.memory, &config.time, &config.threads)
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

	config.keyLen = uint32(len(hash))
	comparisonHash := argon2.IDKey([]byte(password), salt, config.time, config.memory, config.threads, config.keyLen)
	return subtle.ConstantTimeCompare(decodedHash, comparisonHash) == 1, nil
}
