package entities

import (
	"github.com/google/uuid"
	"time"
)

type Base struct {
	ID        uuid.UUID `json:"_id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (b *Base) GenerateID() {
	b.ID = uuid.New()
}

func (b *Base) SetCreatedAt() {
	b.CreatedAt = time.Now()
}

func (b *Base) SetUpdatedAt() {
	b.UpdatedAt = time.Now()
}
