package campaign

import "gorm.io/gorm"

type Repository interface {
	// create contract
	FindAll([]Campaign, error)
	FindByUserID(userID int) ([]Campaign, error)
}

type repository struct {
	// definisikan struct private
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign
	
	err := r.db.Preload("CampaignImage","campaigns_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByUserID(userID int) ([]Campaign, error){
	var campaigns []Campaign
	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages","campaigns_images.is_primary=1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err 
	}

	return campaigns, nil 
}