package repository

import (
	"be13/project/features/feedback"

	"gorm.io/gorm"
)

type feedbackRepository struct {
	db *gorm.DB
}

func NewFeedback(db *gorm.DB) feedback.RepositoryInterface {
	return &feedbackRepository{
		db: db,
	}
}

// AddFeedback implements feedback.RepositoryInterface
func (repo *feedbackRepository) AddFeedback(input feedback.Core) error {
	userGorm := FromUserCore(input) //dari gorm model ke user core yang ada di entities

	tx := repo.db.Create(&userGorm) // proses insert data

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// DeleteFeedback implements feedback.RepositoryInterface
// func (*feedbackRepository) DeleteFeedback(id int) (feedback.Core, error) {
// 	panic("unimplemented")
// }

// // UpdateFeedback implements feedback.RepositoryInterface
// func (*feedbackRepository) UpdateFeedback(id int, input feedback.Core) error {
// 	panic("unimplemented")
// }
