package domain

type User struct { // не йде в бд
	Id         uint64
	FirstName  string
	SecondName string
	Phone      string
	Email      string
	Password   string
}
