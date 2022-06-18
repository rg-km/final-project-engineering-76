package repository

type User struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Loggedin bool   `db:"loggedin"`
	Token    string `db:"token"`
}
type Product struct {
	ID        int64  `db:"id"`
	BookTitle string `db:"book_title"`
	Writer    string `db:"writer"`
	Publisher string `db:"publisher"`
}
