//go:generate gorunpkg github.com/99designs/gqlgen

package go_graphql_gqlgen

import (
	context "context"
)

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateNewEmployee(ctx context.Context, input NewEmployee) (Employee, error) {
	var empID string
	employee := Employee{
		ID:        empID,
		Name:      input.Name,
		Email:     input.Email,
		Contactno: input.Contactno,
		Position:  input.Position,
	}

	return Employee{employee.ID, employee.Name, employee.Email, employee.Contactno, employee.Position}, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Employees(ctx context.Context) ([]Employee, error) {
	var emps []Employee
	return emps, nil
}
