package handler

import (
	"fmt"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// CreateOpening
type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"` //*bool para identificar if != nil
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	// validation body bad request
	if r.Role == "" &&
		r.Company == "" &&
		r.Link == "" &&
		r.Location == "" &&
		r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("request body is empty of malformed")
	}

	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if r.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}

	if r.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}
	return nil
}

//Update Opening

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"` //*bool para identificar if != nil
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	if r.Role != "" || r.Company != "" || r.Link != "" || r.Location != "" || r.Remote != nil || r.Salary > 0 {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}

// Login
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r *LoginRequest) Validate() error {
	if r.Username == "" && r.Password == "" {
		return fmt.Errorf("request body is empty of malformed")
	}
	if r.Username == "" {
		return errParamIsRequired("username", "string")
	}
	if r.Password == "" {
		return errParamIsRequired("password", "string")
	}

	return nil
}

// Register
type UserRegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Profile  string `json:"profile" binding:"required"`
}

func (r *UserRegisterRequest) Validate() error {
	if r.Username == "" && r.Password == "" && r.Email == "" && r.Profile == "" {
		return fmt.Errorf("request body is empty of malformed")
	}

	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}
	if r.Username == "" {
		return errParamIsRequired("username", "string")
	}
	if r.Password == "" {
		return errParamIsRequired("password", "string")
	}

	if r.Profile == "" {
		return errParamIsRequired("profile", "string")
	}

	return nil
}

// Edit
type UpdateUserRequest struct {
	Email    string `json:"email" binding:"omitempty,email"`
	Username string `json:"username" binding:"omitempty"`
	Password string `json:"password" binding:"omitempty"`
	Profile  string `json:"profile" binding:"omitempty"`
}

func (r *UpdateUserRequest) Validate() error {
	//validar campo profile para ser igual as opções pre definidas
	if r.Email != "" || r.Password != "" || r.Username != "" || r.Profile != "" {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
