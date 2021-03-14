package employee

import "fmt"

type Manager struct {
	name           string
	position       string
	yearsInCompany int
}

func NewManager(name, position string, yearsInCompany int) *Manager {
	return &Manager{
		name:           name,
		position:       position,
		yearsInCompany: yearsInCompany,
	}
}

func (m *Manager) BuildProfile() string {
	return fmt.Sprintf("%s is a %s. They have been at the company for %d years.", m.name, m.position, m.yearsInCompany)
}
