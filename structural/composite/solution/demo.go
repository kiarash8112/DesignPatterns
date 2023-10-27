package main

func main() {
	h1 := HumanResource{}
	h2 := HumanResource{}
	t1 := Truck{}

	resource := []components{}
	resource = append(resource, h1, h2, t1)

	team := Team{resources: resource}
	team.deploy()
}
