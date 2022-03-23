package campaign

import "time"

type Campaign struct {
	ID               int
	UserID           int
	Name             string
	ShortDescription string
	Description      string
	Parks            string
	BackerCount      int
	GoalAmount       int
	CurrentAmmount   int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt 			 time.Time
	CampaignImage 	 []CampaignImage
}
type CampaignImage struct {
	ID int
	CampaignID int
	FileName string
	IsPrimary int
	CreatedAt        time.Time
	UpdatedAt 			 time.Time
}