package menu

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/hemanrnjn/d2h/models"
)

var User models.User
var AllPacks []models.Pack
var AllChannels []models.Channel
var AllServices []models.Service

type SuccessSubs struct {
	PackName                               string
	Months, MonthlyCost, Discount, Balance int
}

func ViewBalance() string {
	return "Current Balance is: " + strconv.Itoa(User.Balance)
}

func Recharge(recVal int) string {
	val := recVal
	User.Balance += val
	return "Recharge completed successfully. Current balance is " + strconv.Itoa(User.Balance)
}

func viewAllServices() {
	fmt.Println("Available Packs for Subscription\n")
	for _, val := range AllPacks {
		fmt.Println(val.Name, " :", strings.Join(GetAllChannelsOfPack(val.Name), ", "), " :", strconv.Itoa(val.Price))
	}
	fmt.Println("\nAvailable Channels for Subscription\n")
	for _, val := range AllChannels {
		fmt.Println(val.Ch, ": ", strconv.Itoa(val.Price))
	}
	fmt.Println("\nAvailable Services for Subscription\n")
	for _, val := range AllServices {
		fmt.Println(val.Name, ": ", strconv.Itoa(val.Price))
	}
}

func GetAllChannelsOfPack(name string) []string {
	var channels []string
	for _, val := range AllPacks {
		if val.Name == name {
			for _, chans := range val.Ch {
				channels = append(channels, chans.Ch)
			}
		}
	}
	return channels
}

func SubscribeBase(pack string, months int) (successSubs SuccessSubs, er error) {
	var discountedPrice, discount int
	if strings.ToLower(pack) == strings.ToLower("s") {
		if User.Subs.Pack == "SILVER" {
			er = errors.New("User has already subscribed to this pack")
			return
		} else if User.Subs.Pack == "GOLD" {
			er = errors.New("User already has subscribed to GOLD pack")
			return
		}
		if months >= 3 {
			discount = int((0.1 * float64(50*months)))
			discountedPrice = 50*months - discount
		} else {
			discountedPrice = 50 * months
		}
		if User.Balance >= discountedPrice {
			User.Balance -= discountedPrice
			User.Subs.Ch = append(User.Subs.Ch, AllPacks[0].Ch...)
			User.Subs.Pack = "SILVER"
			successSubs := SuccessSubs{
				PackName:    "SILVER",
				Months:      months,
				MonthlyCost: 50,
				Discount:    discount,
				Balance:     User.Balance,
			}
			return successSubs, nil
		}
		er = errors.New("User doesn't have enough balance in account. Please recharge first")
		return successSubs, er

	} else if strings.ToLower(pack) == strings.ToLower("g") {
		if User.Subs.Pack == "GOLD" {
			er = errors.New("User has already subscribed to this pack")
			return
		}
		if months >= 3 {
			discount = int((0.1 * float64(100*months)))
			discountedPrice = 100*months - discount
		} else {
			discountedPrice = 100 * months
		}
		if User.Balance >= discountedPrice {
			User.Balance -= discountedPrice
			User.Subs.Ch = append(User.Subs.Ch, AllPacks[1].Ch...)
			User.Subs.Pack = "GOLD"
			successSubs := SuccessSubs{
				PackName:    "GOLD",
				Months:      months,
				MonthlyCost: 100,
				Discount:    discount,
				Balance:     User.Balance,
			}
			return successSubs, nil
		}
		er = errors.New("User doesn't have enough balance in account. Please recharge first")
		return successSubs, er
	}
	er = errors.New("Invalid option")
	return
}

func AddChannels(channels []string) (balance int, er error) {
	var newChannels []models.Channel
	var totalCost int
	for _, x := range channels {
		for _, y := range User.Subs.Ch {
			if x == y.Ch {
				er = errors.New("User already has channel: " + x + " in his subscription")
				return
			}
		}
		price := ChannelPrice(x)
		if price == 0 {
			er = errors.New("Channel does not exist")
			return 0, er
		}
		newChannels = append(newChannels, models.Channel{
			Ch:    x,
			Price: price,
		})
	}
	for _, val := range newChannels {
		totalCost += val.Price
	}
	if User.Balance >= totalCost {
		User.Subs.Ch = append(User.Subs.Ch, newChannels...)
		User.Balance -= totalCost
		return User.Balance, nil
	}
	er = errors.New("User doesn't have enough balance in account. Please recharge first")
	return
}

func ChannelPrice(channel string) int {
	for _, val := range AllChannels {
		if strings.ToLower(val.Ch) == strings.ToLower(channel) {
			return val.Price
		}
	}
	return 0
}

func SubscribeSpecial(svc string) (balance int, er error) {
	var f int
	for _, val := range AllServices {
		if strings.Contains(strings.ToLower(val.Name), svc) {
			if IsUserSubscribed(val.Name) {
				er = errors.New("User is already subscribed to this service")
				return
			}
			if User.Balance >= val.Price {
				User.Subs.Svc = append(User.Subs.Svc, models.Service{
					Name:  val.Name,
					Price: val.Price,
				})
				User.Balance -= val.Price
				balance = User.Balance
				return balance, nil
			}
			er = errors.New("User doesn't have enough balance in account. Please recharge first")
			return
		}
		f = 1
	}
	if f == 1 {
		er = errors.New("Service Does not exist")
		return
	}
	return
}

func IsUserSubscribed(name string) bool {
	for _, val := range User.Subs.Svc {
		if val.Name == name {
			return true
		}
	}
	return false
}

func viewSubscription() {
	if User.Subs.Pack == "" {
		println("Currently Subscribed Packs and Channels: " + strings.Join(GetUserChannels(), " + "))
	} else {
		println("Currently Subscribed Packs and Channels: " + User.Subs.Pack + " + " +
			strings.Join(GetUserChannels(), " + "))
	}
	println("Currently subscribed services: " + strings.Join(GetUserServices(), " + "))
}

func GetUserServices() []string {
	var svcs []string
	for _, val := range User.Subs.Svc {
		svcs = append(svcs, val.Name)
	}
	return svcs
}

func GetUserChannels() []string {
	var channels []string
	for _, val := range User.Subs.Ch {
		channels = append(channels, val.Ch)
	}
	return channels
}

func UpdateInfo(email, phone string) {
	User.Email = email
	User.Phonenumber = phone
}
