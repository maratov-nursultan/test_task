package repository

type Users struct {
	Name  string `bun:"name"`
	Iin   string `bun:"iin"`
	Phone string `bun:"phone"`
}
