package main

import (
	"assignment7/database"
	"assignment7/models"
	"fmt"
)

func main() {

	pg_db := database.Connect()
	defer pg_db.Close()

	schemaCreationErr := database.CreateSchema(pg_db)

	if schemaCreationErr != nil {
		panic(schemaCreationErr)
	}

	var companies []models.Company
	if err := pg_db.Model(&companies).Select(); err != nil {
		fmt.Println("Error querying database: ", err)
		return
	}

	adjList := make(map[uint64][]uint64)
	visited := make(map[uint64]bool)

	for _, val := range companies {
		adjList[val.EmpId] = append(adjList[val.EmpId], val.ManagerId)
	}

	var empId, managerId uint64

	fmt.Println("Enter the empID:")
	fmt.Scanln(&empId)
	fmt.Println("Enter the managerID:")
	fmt.Scanln(&managerId)

	// Inserting new data into the Company table
	newCompany := models.Company{
		EmpId:     empId,
		ManagerId: managerId,
	}

	_, insertErr := pg_db.Model(&newCompany).Insert()
	if insertErr != nil {
		fmt.Println("Error inserting data into the database: ", insertErr)
		return
	}

	adjList[empId] = append(adjList[empId], managerId)

	isCycleDetected := dfsDetectCycle(empId, managerId, adjList, visited)

	if isCycleDetected {
		fmt.Println("Cycle Detected!")
	} else {
		fmt.Println("No cycle found.")
	}

}

func dfsDetectCycle(empID uint64, managerID uint64, adjList map[uint64][]uint64, visited map[uint64]bool) bool {

	if visited[managerID] {
		return true
	}

	visited[empID] = true
	visited[managerID] = true

	for _, manager := range adjList[managerID] {
		if dfsDetectCycle(managerID, manager, adjList, visited) {
			return true
		}
	}

	return false
}
