package service

import entity "loan-calculator/domain/calculator/entity"

func generateContract(amount int64) *entity.Contract {
	var contract *entity.Contract
	contract = new(entity.Contract)
	return contract
}
