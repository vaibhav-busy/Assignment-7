:= range companies {
		adjList[val.EmpId] = append(adjList[val.EmpId], val.ManagerId)
	}