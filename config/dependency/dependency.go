package dependency

import (
	"database/sql"
	"fmt"

	"github.com/Genarodaniel/order-system/config/env"
	repository "github.com/Genarodaniel/order-system/internal/infra/database"
	"github.com/Genarodaniel/order-system/internal/usecase"
)

type Repositories struct {
	Order *repository.OrderRepository
}

type UseCases struct {
	Order *usecase.OrderUseCase
}

var (
	Usecase    UseCases
	Repository Repositories
)

var (
	DB *sql.DB
)

func Load() error {
	teste := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		env.Config.DB.User,
		env.Config.DB.Password,
		env.Config.DB.Name,
		env.Config.DB.Host,
		env.Config.DB.Port,
		// TODO: read from Database object when added
		"disable",
	)
	fmt.Println(teste)
	db, err := sql.Open(env.Config.DB.Driver, fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		env.Config.DB.User,
		env.Config.DB.Password,
		env.Config.DB.Name,
		env.Config.DB.Host,
		env.Config.DB.Port,
		// TODO: read from Database object when added
		"disable",
	))
	if err != nil {
		panic(err)
	}

	loadRepositories(db)
	loadUseCases()
	DB = db

	return nil
}

func loadRepositories(db *sql.DB) {
	Repository.Order = repository.NewOrderRepository(db)
	fmt.Println("aqui")
}

func loadUseCases() {
	Usecase.Order = usecase.NewOrderUseCase(Repository.Order)
}
