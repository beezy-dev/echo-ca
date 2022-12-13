package types

type Repository struct {
	ID			int `JSON:"id"`
	Repository 	string `JSON:"repository"`
	Branch		string `JSON:"branch"`
	Folder		string `JSON:"folder"`
}

