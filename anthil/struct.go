package anthil

import (
	"fmt"
)

type Typeroom string

type Room struct {
	Name      string
	Type_room Typeroom
	// ne sert pas avec le code actuel
	// Nbant     int
	Ifant bool
	Link  []*Room
}

type Anthille struct {
	Anthil              []Room
	Tab_Posibility_path [][]*Room
	Lot_uniqueroom_Path [][][]*Room
	Start_nbant         int
}

type Ant struct {
	Name     string
	Pointeur int
	Path     []*Room
	Is_end   bool
}

func (anthil *Anthille) Add_Room(nameRoom string, type_room Typeroom) bool {
	// a revoir manque verificateur name salle unique
	anthil.Anthil = append(anthil.Anthil, Room{Name: nameRoom, Type_room: type_room})
	return true
}

func (anthil *Anthille) Get_Start() (*Room, bool) {
	for index, room := range anthil.Anthil {
		if room.Type_room == Start {
			return &anthil.Anthil[index], true
		}
	}
	return nil, false
}

func (anthil *Anthille) Get_End() (*Room, bool) {
	for index, room := range anthil.Anthil {
		if room.Type_room == End {
			return &anthil.Anthil[index], true
		}
	}
	return nil, false
}

func (anthil *Anthille) Check_salle() bool {
	_, start := anthil.Get_Start()
	_, end := anthil.Get_End()
	return start && end
}

func (anthil *Anthille) Add_link(tabroom []string) bool {
	if len(tabroom) != 2 || tabroom[0] == tabroom[1] {
		return false
	}
	var room1 *Room
	var room2 *Room
	for index, room := range anthil.Anthil {
		if tabroom[0] == room.Name {
			room1 = &anthil.Anthil[index]
		}
		if tabroom[1] == room.Name {
			room2 = &anthil.Anthil[index]
		}
	}
	if room1 == room2 || room1 == nil || room2 == nil {
		return false
	}
	room1.Link = append(room1.Link, room2)
	room2.Link = append(room2.Link, room1)
	return true
}

func (anthil *Anthille) Research_path() {
	var tabpath [][]*Room
	startroom, _ := anthil.Get_Start()
	check_path := func(path []*Room, roomtesting *Room) bool {
		for _, room := range path {
			if room == roomtesting {
				return false
			}
		}
		return true
	}
	var backtracking func(*Room, []*Room)
	backtracking = func(curentroom *Room, path []*Room) {
		path = append(path, curentroom)
		if curentroom.Type_room == End {
			tabpath = append(tabpath, path)
			return
		}
		for _, room := range curentroom.Link {
			if check_path(path, room) {
				// Attention en manipulant un slice
				// a ne pas faire
				// temppath := []*Room(path[:]...)
				var temppath []*Room
				temppath = append(temppath, path...)
				backtracking(room, temppath)
			}
		}
	}
	backtracking(startroom, []*Room{})
	anthil.Tab_Posibility_path = tabpath
}

func (anthil *Anthille) Sort_Path() {
	for index := 1; index < len(anthil.Tab_Posibility_path); index++ {
		if len(anthil.Tab_Posibility_path[index-1]) > len(anthil.Tab_Posibility_path[index]) {
			temp := anthil.Tab_Posibility_path[index]
			anthil.Tab_Posibility_path[index] = anthil.Tab_Posibility_path[index-1]
			anthil.Tab_Posibility_path[index-1] = temp
			index = 0
		}
	}
}

func (anthil *Anthille) Lot_Path() {
	for index, path := range anthil.Tab_Posibility_path {
		var use_room []*Room
		var templotpath [][]*Room
		templotpath = append(templotpath, path)
		use_room = append(use_room, path...)
		for index2 := index + 1; index2 < len(anthil.Tab_Posibility_path); index2++ {
			if Compare(use_room, anthil.Tab_Posibility_path[index2]) {
				templotpath = append(templotpath, anthil.Tab_Posibility_path[index2])
				use_room = append(use_room, anthil.Tab_Posibility_path[index2]...)
			}
		}
		anthil.Lot_uniqueroom_Path = append(anthil.Lot_uniqueroom_Path, templotpath)
	}
}

func Compare(listroom []*Room, path []*Room) bool {
	for _, room := range path {
		if room.Type_room == Start || room.Type_room == End {
			continue
		}
		for index := range listroom {
			if room == listroom[index] {
				return false
			}
		}
	}
	return true
}

func (anthil *Anthille) Dysplais_path() {
	for index, path := range anthil.Tab_Posibility_path {
		// fmt.Println("path ", index+1, " : ", path)
		fmt.Print("path ", index+1, " : [ ")
		for _, room := range path {
			fmt.Print(room.Name, " ")
		}
		fmt.Println("]")
	}
}

func (anthil *Anthille) Dysplaislot_path() {
	for index, lotpath := range anthil.Lot_uniqueroom_Path {
		fmt.Print("======================================================")
		fmt.Println("lot ", index+1, ":")
		for _, path := range lotpath {
			// fmt.Println("	", path)
			fmt.Print("		[ ")
			for _, room := range path {
				fmt.Print(room.Name, " ")
			}
			fmt.Println("]")
		}
	}
}
