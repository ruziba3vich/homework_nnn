package models

var Skills map[string]int

func AddNewValue(skill string, id int) {
	Skills[skill] = id
}
