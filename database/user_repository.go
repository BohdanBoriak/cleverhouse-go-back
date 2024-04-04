package database

import (
	"cleverhouse-go-back/domain"

	"github.com/upper/db/v4"
)

const UsersTableName = "users"

type user struct { //йде в базу даних
	Id         uint64 `db:"id,omitempty"`
	FirstName  string `db:"first_name"`
	SecondName string `db:"second_name"`
	Phone      string `db:"phone"`
	Email      string `db:"email"`
	Password   string `db:"password"`
}

type UserRepository struct {
	coll db.Collection
	sess db.Session
}

func NewUserRepository(session db.Session) UserRepository {
	return UserRepository{
		coll: session.Collection(UsersTableName),
		sess: session,
	}
}

func (r UserRepository) Save(u domain.User) (domain.User, error) {
	user := r.mapDomainToModel(u)
	err := r.coll.InsertReturning(&user)
	if err != nil {
		return domain.User{}, err
	}
	return r.mapModelToDomain(user), nil
}

// перетворює domain.User -> user (який піде в бд)
func (r UserRepository) mapDomainToModel(u domain.User) user {
	return user{
		Id:         u.Id,
		FirstName:  u.FirstName,
		SecondName: u.SecondName,
		Phone:      u.Phone,
		Email:      u.Email,
		Password:   u.Password,
	}
}

// перетворює user -> domain.User (який піде в бд)
func (r UserRepository) mapModelToDomain(u user) domain.User {
	return domain.User{
		Id:         u.Id,
		FirstName:  u.FirstName,
		SecondName: u.SecondName,
		Phone:      u.Phone,
		Email:      u.Email,
		Password:   u.Password,
	}
}
