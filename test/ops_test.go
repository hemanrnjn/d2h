package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hemanrnjn/d2h/menu"
	"github.com/hemanrnjn/d2h/models"
)

func initTest() {
	menu.User.Balance = 100
}

func TestViewBalance(t *testing.T) {
	initTest()
	x := menu.ViewBalance()
	if x != "Current Balance is: 100" {
		t.Error("ViewBalance() failed")
	}
}

func TestRecharge(t *testing.T) {
	initTest()
	x := menu.Recharge(10)
	if x != "Recharge completed successfully. Current balance is 110" {
		t.Error("Recharge() failed")
	}
}

func TestGetAllChannelsOfPack(t *testing.T) {
	initTest()
	x := menu.GetAllChannelsOfPack("Gold Pack")
	if len(x) == 0 {
		t.Error("GetAllChannelsOfPack() failed")
	}

	y := menu.GetAllChannelsOfPack("Silver Pack")
	if len(y) == 0 {
		t.Error("GetAllChannelsOfPack() failed")
	}

	z := menu.GetAllChannelsOfPack("Bronze Pack")
	if len(z) > 0 {
		t.Error("GetAllChannelsOfPack() failed")
	}
}

func TestSubscribeBase(t *testing.T) {
	initTest()
	t.Run("Silver Pack Insuddicient Balance Negative test", func(t *testing.T) {
		var successSubs menu.SuccessSubs
		successSubs, err := menu.SubscribeBase("s", 3)
		if err.Error() != "User doesn't have enough balance in account. Please recharge first" {
			t.Error("Subscribe Base Pack Re subscribe Failed")
		}

		if successSubs.PackName != "" {
			t.Error("Subscribe Base Pack Failed, Packname Invalid")
		}

		if successSubs.Months != 0 {
			t.Error("Subscribe Base Pack Failed, Months Invalid")
		}

		if successSubs.MonthlyCost != 0 {
			t.Error("Subscribe Base Pack Failed, Monthly Cost Invalid")
		}

		if successSubs.Discount != 0 {
			t.Error("Subscribe Base Pack Failed, Discount Invalid")
		}

		if successSubs.Balance != 0 {
			t.Error("Subscribe Base Pack Failed, Balance Invalid")
		}
	})

	t.Run("Gold Pack Insuddicient Balance Negative test", func(t *testing.T) {
		var successSubs menu.SuccessSubs
		successSubs, err := menu.SubscribeBase("g", 3)
		if err.Error() != "User doesn't have enough balance in account. Please recharge first" {
			t.Error("Subscribe Base Pack Re subscribe Failed")
		}

		if successSubs.PackName != "" {
			t.Error("Subscribe Base Pack Failed, Packname Invalid")
		}

		if successSubs.Months != 0 {
			t.Error("Subscribe Base Pack Failed, Months Invalid")
		}

		if successSubs.MonthlyCost != 0 {
			t.Error("Subscribe Base Pack Failed, Monthly Cost Invalid")
		}

		if successSubs.Discount != 0 {
			t.Error("Subscribe Base Pack Failed, Discount Invalid")
		}

		if successSubs.Balance != 0 {
			t.Error("Subscribe Base Pack Failed, Balance Invalid")
		}
	})

	t.Run("Silver Pack Success test", func(t *testing.T) {
		menu.Recharge(200)
		var successSubs menu.SuccessSubs
		successSubs, err := menu.SubscribeBase("s", 3)
		if err != nil {
			t.Error("Subscribe Base Pack Failed")
		}

		if successSubs.PackName != "SILVER" {
			t.Error("Subscribe Base Pack Failed, Packname Invalid")
		}

		if successSubs.Months != 3 {
			t.Error("Subscribe Base Pack Failed, Months Invalid")
		}

		if successSubs.MonthlyCost != 50 {
			t.Error("Subscribe Base Pack Failed, Monthly Cost Invalid")
		}

		if successSubs.Discount != 15 {
			t.Error("Subscribe Base Pack Failed, Discount Invalid")
		}

		fmt.Println(successSubs.Balance)
		if successSubs.Balance != 165 {
			t.Error("Subscribe Base Pack Failed, Balance Invalid")
		}
	})

	t.Run("Silver Pack Resubscribe Negative test", func(t *testing.T) {
		menu.Recharge(135)
		var successSubs menu.SuccessSubs
		successSubs, err := menu.SubscribeBase("s", 3)
		if err.Error() != "User has already subscribed to this pack" {
			t.Error("Subscribe Base Pack Re subscribe Failed")
		}

		if successSubs.PackName != "" {
			t.Error("Subscribe Base Pack Failed, Packname Invalid")
		}

		if successSubs.Months != 0 {
			t.Error("Subscribe Base Pack Failed, Months Invalid")
		}

		if successSubs.MonthlyCost != 0 {
			t.Error("Subscribe Base Pack Failed, Monthly Cost Invalid")
		}

		if successSubs.Discount != 0 {
			t.Error("Subscribe Base Pack Failed, Discount Invalid")
		}

		if successSubs.Balance != 0 {
			t.Error("Subscribe Base Pack Failed, Balance Invalid")
		}
	})

	t.Run("Gold Pack Success test", func(t *testing.T) {
		var successSubs menu.SuccessSubs
		successSubs, err := menu.SubscribeBase("g", 3)
		if err != nil {
			t.Error("Subscribe Base Pack Failed")
		}

		if successSubs.PackName != "GOLD" {
			t.Error("Subscribe Base Pack Failed, Packname Invalid")
		}

		if successSubs.Months != 3 {
			t.Error("Subscribe Base Pack Failed, Months Invalid")
		}

		if successSubs.MonthlyCost != 100 {
			t.Error("Subscribe Base Pack Failed, Monthly Cost Invalid")
		}

		if successSubs.Discount != 30 {
			t.Error("Subscribe Base Pack Failed, Discount Invalid")
		}

		if successSubs.Balance != 30 {
			t.Error("Subscribe Base Pack Failed, Balance Invalid")
		}
	})

	t.Run("Silver Pack after Gold Pack Negative test", func(t *testing.T) {
		menu.Recharge(200)
		var successSubs menu.SuccessSubs
		successSubs, err := menu.SubscribeBase("s", 3)
		if err.Error() != "User already has subscribed to GOLD pack" {
			t.Error("Subscribe Base Silver Pack Failed")
		}

		if successSubs.PackName != "" {
			t.Error("Subscribe Base Pack Failed, Packname Invalid")
		}

		if successSubs.Months != 0 {
			t.Error("Subscribe Base Pack Failed, Months Invalid")
		}

		if successSubs.MonthlyCost != 0 {
			t.Error("Subscribe Base Pack Failed, Monthly Cost Invalid")
		}

		if successSubs.Discount != 0 {
			t.Error("Subscribe Base Pack Failed, Discount Invalid")
		}

		if successSubs.Balance != 0 {
			t.Error("Subscribe Base Pack Failed, Balance Invalid")
		}
	})

	t.Run("Gold Pack Resubscribe Negative test", func(t *testing.T) {
		menu.Recharge(135)
		var successSubs menu.SuccessSubs
		successSubs, err := menu.SubscribeBase("g", 3)
		if err.Error() != "User has already subscribed to this pack" {
			t.Error("Subscribe Base Pack Re subscribe Failed")
		}

		if successSubs.PackName != "" {
			t.Error("Subscribe Base Pack Failed, Packname Invalid")
		}

		if successSubs.Months != 0 {
			t.Error("Subscribe Base Pack Failed, Months Invalid")
		}

		if successSubs.MonthlyCost != 0 {
			t.Error("Subscribe Base Pack Failed, Monthly Cost Invalid")
		}

		if successSubs.Discount != 0 {
			t.Error("Subscribe Base Pack Failed, Discount Invalid")
		}

		if successSubs.Balance != 0 {
			t.Error("Subscribe Base Pack Failed, Balance Invalid")
		}
	})
}

func TestChannelPrice(t *testing.T) {
	initTest()
	t.Run("Channel Price Positive Test", func(t *testing.T) {
		x := menu.ChannelPrice("Nat Geo")
		if x != 20 {
			t.Error("Test Channel Price Failed, Invalid Price")
		}
	})

	t.Run("Channel Price Negative Test", func(t *testing.T) {
		x := menu.ChannelPrice("Colors")
		if x != 0 {
			t.Error("Test Channel Price Failed, Invalid Price")
		}
	})
}

func TestAddChannels(t *testing.T) {
	t.Run("Add Channel Negative Test", func(t *testing.T) {
		menu.User = models.User{}
		channels := []string{"Nat Geo", "Star Plus"}
		bal, err := menu.AddChannels(channels)

		if err.Error() != "User doesn't have enough balance in account. Please recharge first" {
			t.Error("Add Channel Failed")
		}

		if bal != 0 {
			t.Error("Add Channel Failed, Invalid Balance")
		}
	})

	t.Run("Add Channel Positive Test", func(t *testing.T) {
		initTest()
		channels := []string{"Nat Geo", "Star Plus"}
		bal, err := menu.AddChannels(channels)

		if err != nil {
			t.Error("Add Channel Failed")
		}

		if bal != 60 {
			t.Error("Add Channel Failed, Invalid Balance")
		}
	})

	t.Run("Add Channel Negative Test", func(t *testing.T) {
		channels := []string{"Nat Geo", "Star Plus"}
		bal, err := menu.AddChannels(channels)

		if !strings.Contains(err.Error(), "User already has channel: ") {
			t.Error("Add Channel Failed")
		}

		if bal != 0 {
			t.Error("Add Channel Failed, Invalid Balance")
		}
	})

	t.Run("Add Channel Negative Test", func(t *testing.T) {
		channels := []string{"Colors"}
		bal, err := menu.AddChannels(channels)

		if err.Error() != "Channel does not exist" {
			t.Error("Add Channel Failed")
		}

		if bal != 0 {
			t.Error("Add Channel Failed, Invalid Balance")
		}
	})
}

func TestSubscribeSpecial(t *testing.T) {
	t.Run("Subscribe Special Service Negative Test", func(t *testing.T) {
		initTest()
		bal, err := menu.SubscribeSpecial("english")

		if err.Error() != "User doesn't have enough balance in account. Please recharge first" {
			t.Error("Subscribe Special Service Failed")
		}

		if bal != 0 {
			t.Error("Subscribe Special Service Failed, Invalid Balance")
		}
	})

	t.Run("Subscribe Special Service Positive Test", func(t *testing.T) {
		initTest()
		menu.Recharge(200)
		bal, err := menu.SubscribeSpecial("english")

		if err != nil {
			t.Error("Subscribe Special Service Failed")
		}

		if bal != 100 {
			t.Error("Subscribe Special Service Failed, Invalid Balance")
		}
	})

	t.Run("Subscribe Special Service Negative Test", func(t *testing.T) {
		initTest()
		menu.Recharge(200)
		bal, err := menu.SubscribeSpecial("english")

		if err.Error() != "User is already subscribed to this service" {
			t.Error("Subscribe Special Service Failed")
		}

		if bal != 0 {
			t.Error("Subscribe Special Service Failed, Invalid Balance")
		}
	})

	t.Run("Subscribe Special Service Negative Test", func(t *testing.T) {
		initTest()
		menu.Recharge(200)
		bal, err := menu.SubscribeSpecial("dancing")

		if err.Error() != "Service Does not exist" {
			t.Error("Subscribe Special Service Failed")
		}

		if bal != 0 {
			t.Error("Subscribe Special Service Failed, Invalid Balance")
		}
	})
}

func TestGetUserServices(t *testing.T) {
	menu.User = models.User{}
	t.Run("Get User Services Negative Test", func(t *testing.T) {
		svcs := menu.GetUserServices()
		if len(svcs) > 0 {
			t.Error("Get User Services Failed")
		}
	})

	t.Run("Get User Services Positive Test", func(t *testing.T) {
		initTest()
		menu.Recharge(200)
		bal, err := menu.SubscribeSpecial("english")

		if err != nil {
			t.Error("Get User Services, Subscribe Special Service Failed")
		}

		if bal != 100 {
			t.Error("et User Services, Subscribe Special Service Failed with Invalid Balance")
		}

		svcs := menu.GetUserServices()
		if len(svcs) == 0 {
			t.Error("Get User Services Failed")
		}
	})
}

func TestGetUserChannels(t *testing.T) {
	t.Run("Get User Channels Negative Test", func(t *testing.T) {
		menu.User = models.User{}
		chans := menu.GetUserChannels()
		if len(chans) > 0 {
			t.Error("Get User Channels Failed")
		}
	})

	t.Run("Get User Channles Positive Test", func(t *testing.T) {
		menu.User = models.User{}
		menu.Recharge(200)
		channels := []string{"Nat Geo", "Star Plus"}
		bal, err := menu.AddChannels(channels)

		if err != nil {
			t.Error("Get User Channles, Add Channel Failed")
		}

		fmt.Println(bal)
		if bal != 160 {
			t.Error("Get User Channles, Add Channel Failed with Invalid Balance")
		}

		chans := menu.GetUserChannels()
		if len(chans) == 0 {
			t.Error("Get User Channels Failed")
		}
	})
}
