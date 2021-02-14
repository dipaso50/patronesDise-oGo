package guns

import "fmt"

type Revolver struct{}

func (r Revolver) Shoot() {
	fmt.Println("REVOLVER......................")
}
