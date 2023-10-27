package main

import "fmt"

type components interface {
	deploy()
}

type HumanResource struct {
}

func (HumanResource) deploy() {
	fmt.Println("Deploying a human resource")
}

type Truck struct {
}

func (Truck) deploy() {
	fmt.Println("Deploying a truck")
}

type Team struct {
	resources []components
}

func (t *Team) add(resource components) {
	t.resources = append(t.resources, resource)
}

func (t Team) deploy() {
	for _, resource := range t.resources {
		resource.deploy()
	}
}
