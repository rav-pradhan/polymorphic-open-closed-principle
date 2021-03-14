package employee

import (
	"fmt"
	"strings"
)

type Journalist struct {
	name           string
	position       string
	focuses        []string
	yearsInCompany int
}

func NewJournalist(name, position string, focuses []string, yearsInCompany int) *Journalist {
	return &Journalist{
		name:           name,
		position:       position,
		focuses:        focuses,
		yearsInCompany: yearsInCompany,
	}
}

func (j *Journalist) BuildProfile() string {
	return fmt.Sprintf("%s is a %s who has been at the company for %d years. They write primarily on %s",
		j.name, j.position, j.yearsInCompany, j.buildFocuses())
}

func (j *Journalist) buildFocuses() string {
	var focusesCopy strings.Builder
	for i, focus := range j.focuses {
		if i == len(j.focuses)-1 {
			focusesCopy.WriteString(fmt.Sprintf("and %s.", focus))
			break
		}
		focusesCopy.WriteString(fmt.Sprintf("%s, ", focus))
	}
	return focusesCopy.String()
}
