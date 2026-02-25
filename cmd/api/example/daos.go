package example

type DAO struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Data string `db:"data"`
}
