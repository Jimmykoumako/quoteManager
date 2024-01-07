package models

import (
	"gorm.io/gorm"
	"time"
	"errors"
)


const AdminRole ="admin"
// User represents a user in the system
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique;not null;index"`
	Password  string    `json:"password" gorm:"not null"`
	Quotes    []Quote   `json:"quotes" gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

// BeforeCreate hooks into the GORM lifecycle to set default values before saving
func (u *User) BeforeCreate(tx *gorm.DB) error {
	// Set default values or perform any other logic before creating the user
	return nil
}

// CheckPassword verifies if the provided password matches the stored password
func (u *User) CheckPassword(password string) error {
    if u.Password != password {
        return errors.New("password mismatch")
    }
    return nil
}




// package models

// import (
// 	"encoding/base64"
// 	"errors"
// 	"time"
// 	"golang.org/x/crypto/argon2"
// 	"gorm.io/gorm"
// 	"crypto/rand"
// 	"bytes"
// )

// const AdminRole = "admin"

// const (
// 	timeCost     = 1
// 	memoryCost   = 64 * 1024
// 	parallelism  = 2
// 	saltSize     = 16
// 	hashSize     = 32
// )

// // User represents a user in the system
// type User struct {
// 	ID        uint      `json:"id" gorm:"primaryKey"`
// 	Username  string    `json:"username" gorm:"unique;not null;index"`
// 	Password  string    `json:"-" gorm:"not null"`
// 	Salt      string    `json:"-" gorm:"not null"`
// 	Quotes    []Quote   `json:"quotes" gorm:"foreignKey:UserID"`
// 	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
// 	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
// }

// // BeforeCreate hooks into the GORM lifecycle to hash the user's password before saving
// func (u *User) BeforeCreate(tx *gorm.DB) error {
// 	salt, _ := generateRandomBytes(saltSize)
// 	u.Salt = base64.StdEncoding.EncodeToString(salt)
// 	u.Password, _ = HashPasswordWithArgon2(u.Password, salt)
// 	return nil
// }

// // CheckPassword verifies if the provided password matches the hashed password
// func (u *User) CheckPassword(password string) error {
// 	// Decode the stored hashed password string and salt
// 	storedHashedPassword, err := base64.StdEncoding.DecodeString(u.Password)
// 	if err != nil {
// 		return err
// 	}
// 	salt, err := base64.StdEncoding.DecodeString(u.Salt)
// 	if err != nil {
// 		return err
// 	}

// 	// Hash the provided password using Argon2 with the stored salt and parameters
// 	hash := argon2.IDKey([]byte(password), salt, timeCost, memoryCost, parallelism, hashSize)

// 	// Compare the two hashes
// 	if !bytes.Equal(storedHashedPassword, hash) {
// 		return errors.New("incorrect username or password") // Avoid revealing too much information
// 	}

// 	return nil
// }

// // HashPasswordWithArgon2 hashes a password using Argon2
// func HashPasswordWithArgon2(password string, salt []byte) (string, error) {
// 	// Parameters for Argon2 hashing
// 	cost := 12
// 	memory := 64 * 1024
// 	parallelism := 4
// 	keyLength := 32

// 	// Hash the password using Argon2
// 	hash := argon2.IDKey([]byte(password), salt, uint32(cost), uint32(memory), uint8(parallelism), uint32(keyLength))

// 	// Encode the hashed password to base64 before saving to the database
// 	encodedPassword := base64.StdEncoding.EncodeToString(hash)
// 	// Save encodedPassword to the database

// 	return encodedPassword, nil
// }

// // generateRandomBytes generates random bytes of the given length
// func generateRandomBytes(length int) ([]byte, error) {
// 	bytes := make([]byte, length)
// 	_, err := rand.Read(bytes)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return bytes, nil
// }
