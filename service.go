package main

// Service implements the Use Cases layer
type Service interface {
	Insert(c *Customer) (int, error) // returns the ID
	Find(id int) (*Customer, error)
	FindByName(name string) ([]*Customer, error) // returns all names that match the given string
	FindAll() ([]*Customer, error)
	Update(c *Customer) error
	Delete(c *Customer) error
	// SendWarning(c *Customer, u *User) //TODO: ainda não sei como implementar isso
}
