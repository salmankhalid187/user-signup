package usersignup

import (
	"context"

	"github.com/salmankhalid187/user-signup/gen/restapi/operations/signup"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type User struct {
	gorm.Model
	Name string
	Age  string
}

type Repository interface {
	SignupUser(ctx context.Context, params *signup.SignUpUserParams) (bool, error)
}

type repository struct {
	db *gorm.DB
}

// NewRepository pass db reference coming from orm here
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) SignupUser(ctx context.Context, params *signup.SignUpUserParams) (bool, error) {

	r.db.AutoMigrate(&User{})
	// Create
	r.db.Create(&User{Name: *params.Name, Age: *params.Age})
	return true, nil
}
