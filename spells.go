package dnd

import "github.com/brittonhayes/dnd/models"

var _ Spells = &spellsService{}

// The Spells interface shows all of the
// available methods for the spells endpoint
type Spells interface {
	ListSpells() (*models.APIReference, error)
	FindSpell(name string) (*models.Spells, error)
}

type spellsService struct{}

// ListSpells lists the available spells endpoints
func (s *spellsService) ListSpells() (*models.APIReference, error) {
	panic("implement me")
}

// FindSpell searches a specific spell by name
func (s *spellsService) FindSpell(name string) (*models.Spells, error) {
	panic("implement me")
}
