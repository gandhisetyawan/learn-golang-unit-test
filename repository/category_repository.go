package repository

import "learn-golang-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
