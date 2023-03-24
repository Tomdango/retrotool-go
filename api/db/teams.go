package db

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Team struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;" json:"id"`
	Name      string     `json:"name"`
	Owner     string     `json:"owner"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}

func (t *Team) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	t.ID = uuid
	return nil
}

type TeamsRepository struct {
	Initialised bool
	DB          *gorm.DB
}

func (t *TeamsRepository) Initialise() {
	t.DB.AutoMigrate(&Team{})
	t.Initialised = true
}

func (t *TeamsRepository) GetTeamByID(teamID string) *Team {
	if !t.Initialised {
		t.Initialise()
	}

	team := &Team{}
	t.DB.First(team, "id = ?", teamID)

	return team

}

func (t *TeamsRepository) CreateTeam(newTeam *Team) error {
	if !t.Initialised {
		t.Initialise()
	}

	if err := t.DB.Create(newTeam).Error; err != nil {
		return err
	}

	return nil
}

func (t *TeamsRepository) GetAllByOwnerID(ownerID string) *[]Team {
	if !t.Initialised {
		t.Initialise()
	}

	var teams *[]Team
	t.DB.Where("owner = ?", ownerID).Find(&teams)

	return teams
}
