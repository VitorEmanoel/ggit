package git

type Repository struct {
	Path 		string
}

func NewRepository(path string) *Repository{
	var repository = Repository{Path: path}
	return &repository
}


