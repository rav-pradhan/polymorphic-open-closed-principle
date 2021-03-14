package main

import (
	. "github.com/rav-pradhan/personal/open-closed-polymorphic-structs/observes/employee"
	"github.com/rav-pradhan/personal/open-closed-polymorphic-structs/observes/presenter"
)

func main() {
	adoringFan := NewManager("Adoring Fan",
		"Hiring manager",
		5)
	tifa := NewDeveloper("Tifa Lockhart",
		"Senior Developer",
		"Go",
		3)
	bob := NewJournalist("Bob Page",
		"Freelance Journalist",
		[]string{"economics", "trade", "poverty", "social policy"},
		2)

	cmdPresenter := presenter.NewCmdEmployeePresenter(adoringFan, tifa, bob)
	cmdPresenter.DisplayProfiles()
}
