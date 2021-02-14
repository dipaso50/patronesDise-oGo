package main

import (
	"patrones/strategy/player"
	"patrones/strategy/player/guns"
)

func main() {

	player := player.NewPlayer(nil)

	player.Shoot()

	player.ChangeWeapon(guns.Escopeta{})

	player.Shoot()

	player.ChangeWeapon(guns.Revolver{})

	player.Shoot()

}
