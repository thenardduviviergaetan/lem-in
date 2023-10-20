package anthil

import (
	"fmt"
	"strconv"
)

func (anthil *Anthille) Resolve() {
	var Printfinal string
	var nb_line_final int
	for _, lot := range anthil.Lot_uniqueroom_Path {
		tabAnt := Init_tab_Ant(lot, anthil.Start_nbant)
		tempstr := ""
		compt := 0
		for !Is_ant_not_exit(tabAnt) {
			compt++
			tab_path_used := [][2]*Room{}
			for index := range tabAnt {
				tempstr += tabAnt[index].Next_Salle(&tab_path_used)
			}
			tempstr += "\n"
		}
		if nb_line_final == 0 || nb_line_final > compt {
			nb_line_final = compt
			Printfinal = tempstr
		}
	}
	fmt.Print(Printfinal)
}

func (ant *Ant) Next_Salle(tab_path_used *[][2]*Room) string {
	if ant.Is_end {
		return ""
	}
	curentpath := [2]*Room{ant.Path[ant.Pointeur], ant.Path[ant.Pointeur+1]}
	if !path_Not_used(tab_path_used, curentpath) || ant.Path[ant.Pointeur+1].Ifant {
		return ""
	}
	ant.Path[ant.Pointeur].Ifant = false
	ant.Pointeur++
	if ant.Path[ant.Pointeur].Type_room == End {
		ant.Is_end = true
	} else {
		ant.Path[ant.Pointeur].Ifant = true
	}
	*tab_path_used = append(*tab_path_used, curentpath)
	return fmt.Sprint(ant.Name+ant.Path[ant.Pointeur].Name, " ")
}

func path_Not_used(tab_path_used *[][2]*Room, curentpath [2]*Room) bool {
	for _, used_path := range *tab_path_used {
		if used_path == curentpath {
			return false
		}
	}
	return true
}

func Is_ant_not_exit(tab_ant []Ant) bool {
	for _, ant := range tab_ant {
		if !ant.Is_end {
			return false
		}
	}
	return true
}

func Init_tab_Ant(tab_path [][]*Room, nb_ant int) []Ant {
	var tab_ant []Ant
	for id := 1; id <= nb_ant; id++ {
		var ant Ant
		ant.Name = "L" + strconv.Itoa(id) + "-"
		tab_ant = append(tab_ant, ant)
	}
	lengthmin := len(tab_path[0])
	pointer := 0
	for !ant_Not_Path(tab_ant) {
		for _, path := range tab_path {
			if len(path) <= lengthmin && pointer < len(tab_ant) {
				tab_ant[pointer].Path = path
				pointer++
			}
		}
		lengthmin++
	}
	// for _, ant := range tab_ant {
	// 	fmt.Println(ant.Name, ant.Path)
	// }
	return tab_ant
}

func ant_Not_Path(tab_ant []Ant) bool {
	for _, ant := range tab_ant {
		if ant.Path == nil {
			return false
		}
	}
	return true
}
