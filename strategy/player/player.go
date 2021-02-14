package player

import "fmt"

type Weapon interface {
	Shoot()
}
type Player struct {
	weapon *Weapon
}

func NewPlayer(w *Weapon) Player {
	return Player{
		weapon: w,
	}
}

func (p *Player) Shoot() {

	if p.weapon == nil {
		fmt.Println("Player sin armas")
		return
	}

	(*p.weapon).Shoot()

}

func (p *Player) ChangeWeapon(w Weapon) {
	p.weapon = &w
}
