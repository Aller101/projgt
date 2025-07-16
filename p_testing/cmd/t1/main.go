package main

import "fmt"

type User struct {
	ID   int
	Name string
}

type UserRepository interface {
	GetUser(id int) (*User, error)
	SaveUser(user *User) error
}

type PostgresUserRepo struct{}

func (r *PostgresUserRepo) GetUser(id int) (*User, error) {
	fmt.Println("Получение пользователя из PostgreSQL")
	return &User{ID: id, Name: "John Doe"}, nil
}

func (r *PostgresUserRepo) SaveUser(user *User) error {
	fmt.Printf("Сохранение пользователя %v в PostgreSQL\n", user)
	return nil
}

func main() {

	serv := PostgresUserRepo{}

	serv.SaveUser(&User{ID: 1, Name: "Serg"})
	serv.GetUser(1)

	s := Sum(2, 3)
	fmt.Println(s)
}

func Sum(a, b int) int {
	return a + b
}
