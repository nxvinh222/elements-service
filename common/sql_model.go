package common

import "time"

type SQLModel struct {
	Id        int        `json:"-" gorm:"column:id;"`
	FakeId    *UID       `json:"id" gorm:"-"`
	Status    int        `json:"status" gorm:"column:status;default:1"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"updated_at"`
}

func (sm *SQLModel) GenUID(dbType int) {
	fakeUID := NewUID(uint32(sm.Id), dbType, 1)
	sm.FakeId = &fakeUID
}
