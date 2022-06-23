package repository

type User struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Role     string `db:"role"`
	Token    string `db:"token"`
}
type Book struct {
	//ID        int64  `db:"id"`
	BookTitle string `db:"title"`
	Writer    string `db:"writer"`
	Tahun     string `db:"tahun"`
	Kategori  string `db:"kategori"`
	Link      string `db:"link"`
}
