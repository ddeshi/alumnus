package information

type Information struct {
	Title      string
	CreateDate string
	Details    string
	Publiser   Administrator
}

type Administrator struct {
	Name string
	ID   int
}
