package database

import (
	"database/sql"
	"testing"

	"github.com/felipemnz/go-expert-challenge-cleanarch/internal/entity"
	"github.com/stretchr/testify/suite"

	_ "github.com/mattn/go-sqlite3"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db   *sql.DB
	repo entity.OrderRepositoryInterface
}

func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
	suite.repo = NewOrderRepository(suite.Db)
}

func (suite *OrderRepositoryTestSuite) SetupTest() {
	suite.Db.Exec("DELETE FROM orders;")
}

func (suite *OrderRepositoryTestSuite) TearDownSuite() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGivenAnOrder_WhenSave_ThenShouldSaveOrder() {
	order := suite.persistOrder("123", 10.0, 5.0)

	var orderResult entity.Order
	err := suite.Db.QueryRow("Select id, price, tax, final_price from orders where id = ?", order.ID).
		Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)

	suite.NoError(err)
	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}

func (suite *OrderRepositoryTestSuite) TestListOrder_whenHasRegisters() {
	suite.persistOrder("456", 10.0, 5.0)
	suite.persistOrder("789", 12.0, 7.0)
	suite.persistOrder("101112", 14.0, 9.0)

	orders, err := suite.repo.List()
	suite.NoError(err)
	suite.Equal(3, len(orders))
	suite.Equal("456", orders[0].ID)
	suite.Equal("789", orders[1].ID)
	suite.Equal("101112", orders[2].ID)
}

func (suite *OrderRepositoryTestSuite) persistOrder(id string, price float64, tax float64) entity.Order {
	order, err := entity.NewOrder(id, price, tax)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	err = suite.repo.Save(order)
	suite.NoError(err)

	return *order
}
