package main

import (
	ant "lem-in/anthil"
)

func main() {
	anthil, check := ant.Init_Data()
	if !check {
		return
	}
	// fmt.Println(anthil)
	anthil.Research_path()
	// anthil.Dysplais_path()
	anthil.Sort_Path()
	// anthil.Dysplais_path()
	anthil.Lot_Path()
	// anthil.Dysplaislot_path()
	// anthil.Resolve()
}
