// Package updating implements service and repository to update data
package updating

import "github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/entity"

// Service implemets methods to updating
type Service interface {
	ChangePassword(u entity.User) error
}

// Repository implements methods to repository
type Repository interface {
	UpdateUser(u entity.User) error
}

type service struct {
	bR Repository
}

// NewService implements a new service to updating
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) ChangePassword(u entity.User) error {
	b, err := entity.Bcrypt(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(b)
	return s.bR.UpdateUser(u)
}
