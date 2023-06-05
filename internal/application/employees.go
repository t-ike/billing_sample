package application

type EmployeeService struct {
}

func (s *EmployeeService) PaySalary(payday int32) {
	employees := EmployeeRepo.GetEmployeesByPayDay(payday)

	for _, employee := range employees {
		// すでに今月支払い済みの場合は飛ばす
		if s.employeeService.AlreadyPayed(employee, payday) {
			continue
		}

		s.uow.Do(
			repos.SaralyLog.Create()
			s.payService.pay(employee.id, payday, ...)
		)
	}

}
