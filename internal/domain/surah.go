package domain

type Surah struct {
	ID          int64
	Number      int
	NameArabic  string
	NameEnglish string
	NameRussian string
	AyahCount   int
}
