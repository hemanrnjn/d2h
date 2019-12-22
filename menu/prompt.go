package menu

import (
	"fmt"
	"strconv"
	"strings"
)

func promptViewBalance() {
	bal := ViewBalance()
	println(bal)
}

func promptRecharge() {
	var recVal int
	print("Enter the amount to recharge: ")
	fmt.Scan(&recVal)
	rechrgVal := Recharge(recVal)
	println(rechrgVal)
}

func promptSubscribeBase() {
	var pack string
	var months int
	var successSubs SuccessSubs
	print("Enter the Pack you wish to subscribe: (Silver: ‘S’, Gold: ‘G’): ")
	fmt.Scanln(&pack)
	print("Enter the months: ")
	fmt.Scanln(&months)
	successSubs, err := SubscribeBase(pack, months)
	if err != nil {
		println(err.Error())
		return
	}
	successSubscription(successSubs.PackName, successSubs.Months, successSubs.MonthlyCost,
		successSubs.Discount, successSubs.Balance)
}

func successSubscription(pack string, months, price, discount, balance int) {
	println("You have successfully subscribed the following packs - " + pack)
	println("Monthly price: " + strconv.Itoa(price) + " Rs.")
	println("No of months: " + strconv.Itoa(months))
	println("Subscription Amount: " + strconv.Itoa(price*months) + " Rs.")
	println("Discount applied: " + strconv.Itoa(discount) + " Rs.")
	println("Final Price after discount: " + strconv.Itoa((price*months)-discount) + " Rs.")
	println("Account balance: " + strconv.Itoa(balance) + " Rs.")
	println("Email notification sent successfully")
	println("SMS notification sent successfully")
}

func promptAddChannels() {
	print("Enter channel names to add (separated by commas): ")
	scanner.Scan()
	channels := strings.Split(scanner.Text(), ",")
	balance, err := AddChannels(channels)
	if err != nil {
		println(err.Error())
		return
	}
	println("Channels added successfully.")
	println("Account Balance: " + strconv.Itoa(balance))
}

func promptSubscribeSpecial() {
	print("Enter the service name: ")
	scanner.Scan()
	text := scanner.Text()
	svc := strings.ToLower(text)
	balance, err := SubscribeSpecial(svc)
	if err != nil {
		println(err.Error())
		return
	}
	println("Service subscribed successfully")
	println("Account balance: " + strconv.Itoa(balance) + " Rs.")
	println("Email notification sent successfully")
	println("SMS notification sent successfully")
}

func promptUpdateInfo() {
	print("Enter the email: ")
	scanner.Scan()
	email := scanner.Text()
	print("Enter Phone: ")
	scanner.Scan()
	phone := scanner.Text()
	UpdateInfo(email, phone)
	println("Email and Phone updated successfully")
}

func println(val string) {
	fmt.Println(val)
}
