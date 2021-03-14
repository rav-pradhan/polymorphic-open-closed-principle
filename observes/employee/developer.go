package employee

import "fmt"

type Developer struct {
	name                       string
	primaryProgrammingLanguage string
	position                   string
	yearsInCompany             int
}

func NewDeveloper(name, position, primaryProgrammingLanguage string, yearsInCompany int) *Developer {
	return &Developer{
		name:                       name,
		primaryProgrammingLanguage: primaryProgrammingLanguage,
		position:                   position,
		yearsInCompany:             yearsInCompany,
	}
}

func (d *Developer) BuildProfile() string {
	return fmt.Sprintf("%s is a %s. They have been at the company for %d years. They primarily code in %s.",
		d.name, d.position, d.yearsInCompany, d.primaryProgrammingLanguage)
}
