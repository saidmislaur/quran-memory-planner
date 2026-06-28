package domain

type Word struct {
	ID       int64
	SurahID  int64
	AyahID   int64
	Arabic   string
	Root     string
	Position int
}
