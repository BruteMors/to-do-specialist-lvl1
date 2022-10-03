package main

import (
	"fmt"
	"to-do-specialist-lvl1/internal/controller/stdio"
	"to-do-specialist-lvl1/internal/database"
	"to-do-specialist-lvl1/internal/domain/service"
)

func main() {
	var str string
	_, err := fmt.Scan(&str)
	if err != nil {
		return
	}
	if str == "StartApp" {
		fmt.Println("App is started")
		storage := database.NewDayToDoStorage()
		srvc := service.NewDayToDoService(storage)
		handler := stdio.NewDayToDoHandler(srvc)

		for {
			cmd, err := handler.GetCommand()
			if err != nil {
				fmt.Println(err)
				continue
			}
			isExit, err := handler.ParseCommand(cmd)
			if err != nil {
				fmt.Println(err)
			}
			if isExit {
				fmt.Println("App is terminated")
				return
			}
		}
	}
}
