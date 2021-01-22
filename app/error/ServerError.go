package error

import "fmt"

type ServerError struct {
	Code int
	Desc string
	error
}

func (s *ServerError) Error() string {
	return fmt.Sprintf("Server code %d: %s", s.Code, s.Desc)
}
