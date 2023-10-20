package anthil

import (
	"os"
	"strconv"
	"strings"

	"lem-in/utils"
)

func Init_Data() (Anthille, bool) {
	arg := os.Args[1:]
	if len(arg) != 1 {
		utils.Print_err(utils.Err{Message: "Veuillez entrer un nombre d'argument valide"})
		return Anthille{}, false
	}
	tabfile := utils.Readfile(arg[0])
	var anthil Anthille
	var err error
	tabline := strings.Split(string(tabfile), "\n")
	anthil.Start_nbant, err = strconv.Atoi(tabline[0])
	utils.Check_err(err)
	if anthil.Start_nbant <= 0 {
		utils.Print_err(utils.Err{Message: "ERROR: invalid data format"})
		return anthil, false
	}
	for index := 1; index < len(tabline); index++ {
		switch tabline[index] {
		case "##start":
			index++
			anthil.Add_Room(strings.Split(tabline[index], " ")[0], Start)
			continue
		case "##end":
			index++
			anthil.Add_Room(strings.Split(tabline[index], " ")[0], End)
			continue
		}
		ispath := strings.Split(tabline[index], "-")
		isroominfo := strings.Split(tabline[index], " ")
		switch {
		case len(ispath) == 2:
			if !anthil.Add_link(ispath) {
				utils.Print_err(utils.Err{
					Message: "ERROR: invalid data format",
					Reason:  tabline[index],
				})
				return anthil, false
			}
		case len(isroominfo) == 3:
			if !anthil.Add_Room(isroominfo[0], Standard) {
				utils.Print_err(utils.Err{
					Message: "ERROR: invalid data format",
					Reason:  tabline[index],
				})
				return anthil, false
			}
		case tabline[index][0] == '#':
			continue
		default:
			utils.Print_err(utils.Err{
				Message: "ERROR: invalid data format",
				Reason:  tabline[index],
			})
			return anthil, false
		}
	}
	if !anthil.Check_salle() {
		utils.Print_err(utils.Err{
			Message: "ERROR: invalid data format",
		})
		return anthil, false
	}
	return anthil, true
}
