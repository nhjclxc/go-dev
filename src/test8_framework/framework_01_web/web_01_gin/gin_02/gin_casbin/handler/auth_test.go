package handler

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {

	fmt.Println(encryption("superadmin"))
	fmt.Println(encryption("admin"))
	fmt.Println(encryption("common"))

	//$2a$10$K5b1SMcqqObPWyT5dCVKauepxwsAHLJCruRam3oNfg5SXArFeCojq <nil>
	//$2a$10$HMv7fwShhxKUCPyvDUb.ou8xNjIIQSeOqhHpscu5It.n4R46jjHQK <nil>
	//$2a$10$H5iz6nT6BHBDuBYfBMrR0.lWRbhTpyYBqFqyS9B4W/TQttH9Ghmca <nil>

}
