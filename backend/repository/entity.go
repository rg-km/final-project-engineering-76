package repository

type User struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Loggedin bool   `db:"loggedin"`
	Token    string `db:"token"`
}
type Book struct {
	//ID        int64  `db:"id"`
	BookTitle string `db:"title"`
	Writer    string `db:"writer"`
	Publisher string `db:"publisher"`
}
