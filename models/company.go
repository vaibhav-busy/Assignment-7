package models

type Company struct {
	EmpId uint64 `pg:",pk"`
	ManagerId uint64 `pg:",pk"`
}