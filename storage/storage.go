package storage

import (
	"github.com/jiaqi-yin/go-australian-holidays/domain"
)

type Storage interface {
	Save(holidays []domain.Holiday)
	Load() []domain.Holiday
}
