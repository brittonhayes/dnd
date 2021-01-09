package dnd

import "github.com/brittonhayes/dnd/models"

var _ Spells = &SpellsService{}

// The Spells interface shows all of the
// available methods for the spells endpoint
type Spells interface {
	ListSpells() (*models.APIReference, error)
	FindSpell(name string) (*models.Spells, error)
}

type SpellsService struct{}

// ListSpells lists the available spells endpoints
func (s *SpellsService) ListSpells() (*models.APIReference, error) {
	panic("implement me")
}

// FindSpell searches a specific spell by name
func (s *SpellsService) FindSpell(name string) (*models.Spells, error) {
	panic("implement me")
}
