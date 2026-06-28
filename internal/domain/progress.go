package domain

import "time"

type ReviewLevel string

const (
	ReviewHard   ReviewLevel = "hard"
	ReviewNormal ReviewLevel = "normal"
	ReviewEasy   ReviewLevel = "easy"
)

type WordProgress struct {
	UserID         int64
	WordID         int64
	ReviewCount    int
	EasyCount      int
	LastReviewedAt *time.Time
	NextReviewAt   time.Time
	IsMastered     bool
}

func (p *WordProgress) ApplyReview(level ReviewLevel, now time.Time) {
	p.ReviewCount++
	p.LastReviewedAt = &now

	switch level {
	case ReviewHard:
		p.EasyCount = 0
		p.IsMastered = false
		p.NextReviewAt = now.Add(10 * time.Minute)

	case ReviewNormal:
		p.EasyCount = 0
		p.IsMastered = false
		p.NextReviewAt = now.Add(24 * time.Hour)

	case ReviewEasy:
		p.EasyCount++
		p.NextReviewAt = now.Add(72 * time.Hour)

		if p.EasyCount >= 3 {
			p.IsMastered = true
		}
	}
}
