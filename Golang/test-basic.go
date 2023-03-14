package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Manager struct {
	FullName       string `json:"full_name"`
	Position       string `json:"position"`
	Age            int32  `json:"age"`
	YearsInCompany int32  `json:"years_in_company"`
}

func EncodeManager(manager *Manager) (io.Reader, error) {
	// var res Manager
	if manager.YearsInCompany == 0 {

		type Manager struct {
			FullName string `json:"full_name"`
			Position string `json:"position"`
			Age      int32  `json:"age"`
		}
		res, err := json.Marshal(&Manager{
			FullName: manager.FullName,
			Position: manager.Position,
			Age:      manager.Age,
		})
		if err != nil {
			return nil, err
		}
		result := string(res)
		return strings.NewReader(result), nil
	}
	res, err := json.Marshal(manager)
	if err != nil {
		return nil, err
	}
	result := string(res)
	return strings.NewReader(result), nil
}

func main() {

	fmt.Println(EncodeManager(&Manager{
		FullName:       "Javlonbek",
		Position:       "Backend",
		Age:            17,
		YearsInCompany: 0,
	}))
}
