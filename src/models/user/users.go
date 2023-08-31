package user

import (
	"github.com/WalterSmith9/trainee-assignment/src/helpers/db"
	"github.com/WalterSmith9/trainee-assignment/src/models/entities"
	"github.com/WalterSmith9/trainee-assignment/src/swagger/generated/models"
)

// EditUser Edits a user's segments
func EditUser(args *models.UserEditRequest) error {

	newUserEntity := entities.User{
		UserId: int(*args.UserID),
	}
	_, err := db.GetConnection().Model(&newUserEntity).Insert()
	//if err != nil {
	//	return err
	//}

	err = db.GetConnection().
		Model(&newUserEntity).
		Where("user_id = ?", int(*args.UserID)).
		Select()
	if err != nil {
		return err
	}

	newSegmentEntity := entities.Segment{}
	newRecordEntity := entities.Record{
		UserId: newUserEntity.Id,
	}

	for _, newSegmentEntity.Slug = range args.SlugsAdd {
		err = db.GetConnection().
			Model(&newSegmentEntity).
			Where("slug = ?", newSegmentEntity.Slug).
			Select()
		if err != nil {
			return err
		}

		newRecordEntity.Segment = newSegmentEntity.Id

		_, err = db.GetConnection().Model(&newRecordEntity).Insert()
	}

	for _, newSegmentEntity.Slug = range args.SlugsDelete {
		err = db.GetConnection().
			Model(&newSegmentEntity).
			Where("slug = ?", newSegmentEntity.Slug).
			Select()
		if err != nil {
			return err
		}

		newRecordEntity.Segment = newSegmentEntity.Id

		_, err = db.GetConnection().
			Model(&newRecordEntity).
			Where("user_id = ?", newRecordEntity.UserId).
			Where("segment_id = ?", newRecordEntity.Segment).
			Delete()
		//if err != nil {
		//	return err
		//}
	}

	return nil
}

//// GetUserSegment Edits a user's segments
//func GetUserSegment(args *models.UserSegmentsGetRequest) error {
//
//	newUserSegments := make([]entities.Record, 0)
//
//
//	return nil
//}
