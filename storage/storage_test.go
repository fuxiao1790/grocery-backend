package storage

import (
	"fmt"
	"grocery-backend/dto"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const MONGO_DB_URI = "mongodb://localhost:27272"
const DB_NAME = "GROCERY_DB"

const TEST_STORE_COUNT = 5

var TEST_ITEM = &dto.Item{
	IconUri: "test uri",
	Name:    "test name",
	Price:   10,
	ID:      "619d1f0ac5aa3b27c13861eb",
	StoreID: "619d1f09c5aa3b27c13861ac",
}

var TEST_ORDER = &dto.Order{
	ItemList:  []*dto.OrderItemData{},
	Location:  "test location",
	CreatedAt: 0,
	UserData: &dto.User{
		Username: "",
		ID:       "61a262824270c01dde9b34cb",
	},
	StoreData: &dto.Store{
		Location: "",
		ID:       "619d1f09c5aa3b27c13861ac",
		Name:     "",
	},
	ID: "",
}

var TEST_STORE = &dto.Store{
	Location: "test location",
	Name:     "test store name",
}

func TestMain(m *testing.M) {
	m.Run()
}

func Test_GetItemListWithPriceQuery(t *testing.T) {
	st, err := NewMongoDBStorage(&Config{
		Uri:  MONGO_DB_URI,
		Name: DB_NAME,
	})
	if err != nil {
		t.FailNow()
	}

	res, err := st.GetItemList(0, 100, "61a6392d83692e717b418cf4", &dto.ItemListQuery{Name: "", PriceMax: 140, PriceMin: 100})
	if err != nil {
		t.FailNow()
	}

	for _, item := range res {
		assert.True(t, item.Price > 100)
		assert.True(t, item.Price < 140)
	}
}

func Test_GetItemListWithNameQuery(t *testing.T) {
	st, err := NewMongoDBStorage(&Config{
		Uri:  MONGO_DB_URI,
		Name: DB_NAME,
	})
	if err != nil {
		t.FailNow()
	}

	res, err := st.GetItemList(0, 100, "61a6392d83692e717b418cf4", &dto.ItemListQuery{Name: "name"})
	if err != nil {
		t.FailNow()
	}

	assert.True(t, len(res) > 0)

	for _, item := range res {
		assert.True(t, strings.Contains(item.Name, "name"))
	}
}

func Test_CreateOrder(t *testing.T) {
	st, err := NewMongoDBStorage(&Config{
		Uri:  MONGO_DB_URI,
		Name: DB_NAME,
	})
	if err != nil {
		t.FailNow()
	}

	err = st.CreateOrder(
		map[string]int{TEST_ITEM.ID: 3},
		TEST_ORDER.Location,
		TEST_ORDER.StoreData.ID,
		TEST_ORDER.UserData.ID,
	)
	assert.Nil(t, err)
}

func Test_User(t *testing.T) {
	st, err := NewMongoDBStorage(&Config{
		Uri:  MONGO_DB_URI,
		Name: DB_NAME,
	})
	if err != nil {
		t.FailNow()
	}

	{
		err := st.CreateUser(&User{
			Username:       "test-user-name",
			HashedPassword: "test-hashed-password",
		})
		if err != nil {
			t.FailNow()
		}
	}

	{
		user, err := st.GetUser(&User{
			Username: "test-user-name",
		})
		if err != nil {
			t.FailNow()
		}

		assert.NotNil(t, user)
	}
}

func Test_CreateTestData(t *testing.T) {
	st, err := NewMongoDBStorage(&Config{
		Uri:  MONGO_DB_URI,
		Name: DB_NAME,
	})
	if err != nil {
		t.FailNow()
	}

	createTestData(st)
}

func createTestData(st Storage) {
	for i := 0; i < TEST_STORE_COUNT; i++ {
		st.CreateStore(&dto.Store{
			Location: fmt.Sprintf("location %d", i),
			Name:     fmt.Sprintf("name %d", i),
			ID:       "",
		})
	}

	storeList, _ := st.GetStoreList(0, TEST_STORE_COUNT)

	for _, store := range storeList {
		for i := 0; i < 15; i++ {
			st.CreateItem(&dto.Item{
				IconUri: fmt.Sprintf("%s | uri %d", store.Name, i),
				Name:    fmt.Sprintf("%s | name %d", store.Name, i),
				Price:   i * 10,
				ID:      "",
				StoreID: store.ID,
			})
		}
	}
}
