package repositories

import "ettoreMB/controle_gastos/internal/entities"

type ExpenseRepository interface {
	Create(item entities.Expense) error
	GetByID(id string) (*entities.Expense, error)
	GetAll() ([]entities.Expense, error)
	Update(entities.Expense) error
	Delete(id string) error
}
