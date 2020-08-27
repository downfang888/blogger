package model

type Auth struct {
	Id           int64     `db:"id"`
	Username   string     `db:"username"`
	Password      string    `db:"password"`
}
