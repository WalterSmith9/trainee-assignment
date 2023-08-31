package segment

import (
	_ "encoding/json"
	"errors"
	"github.com/WalterSmith9/trainee-assignment/src/helpers/db"
	"github.com/WalterSmith9/trainee-assignment/src/models/entities"
	"github.com/WalterSmith9/trainee-assignment/src/swagger/generated/models"
)

// AddSegment Adds a new segment
func AddSegment(args *models.SegmentCreateRequest) error {

	newSegmentEntity := entities.Segment{
		Slug: *args.Slug,
	}

	_, err := db.GetConnection().Model(&newSegmentEntity).Insert()
	if err != nil {
		return err
	}

	return nil
}

// DeleteSegment Removes a existing segment
func DeleteSegment(args *models.SegmentDeleteRequest) error {

	deleteSegmentEntity := entities.Segment{
		Slug: *args.Slug,
	}

	res, err := db.GetConnection().
		Model(&deleteSegmentEntity).
		Where("slug = ?", deleteSegmentEntity.Slug).
		Delete()
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return errors.New("Unknown segment slug! ")
	}

	return nil
}
