package services

import (
	"app/server_side/models"
	"errors"

	"github.com/astaxie/beego/logs"
)

// PostFavoriteJob ...
func PostFavoriteJob(favoriteJob models.FavoriteJob) (favoriteJobID int64, err error) {
	tx := db.Begin()
	var currentFavoriteJob []*models.FavoriteJob
	currentFavoriteJob, err = models.FindFavoriteJobWithUserIDAndJobID(tx, favoriteJob.UserID, favoriteJob.JobID)
	if len(currentFavoriteJob) != 0 {
		tx.Rollback()
		return 0, errors.New("この案件はすでに気になる登録されています。")
	}
	favoriteJobID, err = models.CreateFavoriteJob(tx, favoriteJob)
	if err != nil {
		logs.Error(err)
		return
	}
	err = tx.Commit().Error
	return
}

// DeleteFavoriteJobWithUserIDAndJobID ...
func DeleteFavoriteJobWithUserIDAndJobID(favoriteJob models.FavoriteJob) (err error) {
	tx := db.Begin()
	var currentFavoriteJob []*models.FavoriteJob
	currentFavoriteJob, err = models.FindFavoriteJobWithUserIDAndJobID(tx, favoriteJob.UserID, favoriteJob.JobID)
	if err != nil {
		logs.Error(err)
		return
	}
	if len(currentFavoriteJob) != 0 {
		err = models.DeleteFavoriteJob(tx, favoriteJob.UserID, favoriteJob.JobID)
		if err != nil {
			logs.Error(err)
			return
		}
	} else {
		tx.Rollback()
		return errors.New("この案件はすでに気になる登録から削除されています。")
	}

	err = tx.Commit().Error
	return
}
