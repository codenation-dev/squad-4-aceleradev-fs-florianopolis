package customer

import "codenation/squad-4-aceleradev-fs-florianopolis/models"

// Repository repository interface
type Repository interface {
	Create(c *models.Customer) (*models.Customer, error) // returns the ID
	Read(id int) (*models.Customer, error)
	ReadByName(pattern string) ([]*models.Customer, error) // returns all names that match the given string
	ReadAll() ([]*models.Customer, error)
	Update(c *models.Customer) error
	Delete(id int) error
	// SendWarning(c *models.Customer, u *User) //TODO: ainda não sei como implementar isso
}
