import "github.com/thanhtuan260593/todo/models"

func (db *Database) ListAllItems() ([]models.Item, error) {
	var items []models.Item
	if err := db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}