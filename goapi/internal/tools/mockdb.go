package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username: "alex",
	},
	"john": {
		AuthToken: "456DEF",
		Username: "john",
	},
	"james": {
		AuthToken: "789GHI",
		Username: "james",
	},
}

var mockCoinDetails = map[string]CoinDetails {
	"alex": {
		Coins: 100,
		Username: "alex",
	},
	"john": {
		Coins: 200,
		Username: "john",
	},
	"james": {
		Coins: 300,
		Username: "james",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB Call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return  nil
	}

	return &clientData
}

func (d *mockDB) GetUsersCoins(username string) *CoinDetails {
	// Simulate DB Call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return  nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}