package presenter

import (
	"fmt"

	"github.com/rav-pradhan/personal/open-closed-polymorphic-structs/observes/employee"
)

type CmdEmployeePresenter struct {
	Employees []employee.Employee
}

func NewCmdEmployeePresenter(employees ...employee.Employee) *CmdEmployeePresenter {
	var employeesList []employee.Employee
	for _, employee := range employees {
		employeesList = append(employeesList, employee)
	}
	return &CmdEmployeePresenter{Employees: employeesList}
}

func (c *CmdEmployeePresenter) DisplayProfiles() {
	for _, employee := range c.Employees {
		fmt.Println(employee.BuildProfile())
	}
}
