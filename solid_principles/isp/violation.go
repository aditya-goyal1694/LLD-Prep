package main

import "fmt"

type Worker interface {
	WriteCode() error
	RunTests() error
	DeployApplication() error
	CreateDesignDocument() error
}

type SoftwareEngineer struct{}

func (s *SoftwareEngineer) WriteCode() error {
	fmt.Println("Writing code...")
	return nil
}

func (s *SoftwareEngineer) RunTests() error {
	fmt.Println("Running tests...")
	return nil
}

func (s *SoftwareEngineer) DeployApplication() error {
	fmt.Println("Deploying application...")
	return nil
}

func (s *SoftwareEngineer) CreateDesignDocument() error {
	fmt.Println("Creating design document...")
	return nil
}

type Intern struct{}

func (i *Intern) WriteCode() error {
	fmt.Println("Writing code...")
	return nil
}

func (i *Intern) RunTests() error {
	fmt.Println("Running tests...")
	return nil
}

func (i *Intern) DeployApplication() error {
	return fmt.Errorf("interns cannot deploy")
}

func (i *Intern) CreateDesignDocument() error {
	return fmt.Errorf("interns do not create design documents")
}

func StartWorkday(w Worker) {
	_ = w.WriteCode()
	_ = w.RunTests()
	_ = w.DeployApplication()
	_ = w.CreateDesignDocument()
}

func main() {
	engineer := &SoftwareEngineer{}
	intern := &Intern{}

	StartWorkday(engineer)
	StartWorkday(intern)
}
