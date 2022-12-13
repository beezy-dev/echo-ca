package types

type Jobs struct { 
	ID		int `JSON:"id"`
	Type 	string `JSON:"type"`
	Start	string `JSON:"start"`
	End 	string `JSON:"end"`
	Success	bool `JSON:"success"`
}