package object

// import (
// 	"fmt"

// 	"github.com/pkg/errors"
// 	"golang.org/x/crypto/bcrypt"
// )

type (
	StatusID    = int64

	// Account account
	Status struct {
		ID StatusID `json:"id"`

		Account *Account `json:"account"`

		AccountId int64 `json:"-" db:"account_id"`

		Content *string `json:"content"`

		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`
	}
)

// // Check if given password is match to account's password
// func (a *Account) CheckPassword(pass string) bool {
// 	return bcrypt.CompareHashAndPassword([]byte(a.PasswordHash), []byte(pass)) == nil
// }

// // Hash password and set it to account object
// func (a *Account) SetPassword(pass string) error {
// 	passwordHash, err := generatePasswordHash(pass)
// 	if err != nil {
// 		return fmt.Errorf("generate error: %w", err)
// 	}
// 	a.PasswordHash = passwordHash
// 	return nil
// }

// func generatePasswordHash(pass string) (PasswordHash, error) {
// 	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
// 	if err != nil {
// 		return "", fmt.Errorf("hashing password failed: %w", errors.WithStack(err))
// 	}
// 	return PasswordHash(hash), nil
// }
