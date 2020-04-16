import "github.com/thanhtuan260593/todo/models"

func (db *Database) CompleteItem(id uint) error {
	var item models.Item
	if err := db.First(&item, id).Error; err != nil {
		return err
	}
	item.IsDone = true
	if err := db.Save(&item).Error; err != nil {
		return err
	}
	return nil
}