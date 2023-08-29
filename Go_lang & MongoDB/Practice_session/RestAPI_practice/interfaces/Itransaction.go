
// interfaces.go
package interfaces

import "Rest-API/models"

type ITransactions interface {
	CreateTransaction(transaction *models.Transactions) (string, error)

	
}
