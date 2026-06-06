package main

import "fmt"

type Developer interface{
    WriteCode() error
}

type Tester interface {
    RunTests() error
}

type Deployer interface {
    DeployApplication() error
}

type Documenter interface {
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

func AssignCodingTask(d Developer) {
	_ = d.WriteCode()
}

func AssignTestingTask(t Tester) {
	_ = t.RunTests()
}

func AssignDeploymentTask(d Deployer) {
	_ = d.DeployApplication()
}

func AssignDocumentationTask(d Documenter) {
	_ = d.CreateDesignDocument()
}

func main() {
	engineer := &SoftwareEngineer{}
	intern := &Intern{}

	AssignCodingTask(engineer)
	AssignTestingTask(engineer)
	AssignDeploymentTask(engineer)
	AssignDocumentationTask(engineer)

	fmt.Println()

	AssignCodingTask(intern)
	AssignTestingTask(intern)
}