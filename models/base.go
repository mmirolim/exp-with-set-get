package models

import "time"

type Base struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	DeletedBy string
	CreatedBy string
	UpdatedBy string
}

func (b *Base) Created(by string, at ...time.Time) {
	b.CreatedBy = by
	if len(at) == 1 {
		b.CreatedAt = at[0]
	} else {
		b.CreatedAt = time.Now()
	}
}

func (b *Base) Updated(by string, at ...time.Time) {
	b.UpdatedBy = by
	if len(at) == 1 {
		b.UpdatedAt = at[0]
	} else {
		b.UpdatedAt = time.Now()
	}
}

func (b *Base) Deleted(by string, at ...time.Time) {
	b.DeletedBy = by
	if len(at) == 1 {
		b.DeletedAt = at[0]
	} else {
		b.DeletedAt = time.Now()
	}
}
