package menu

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func Landing() {
Loop:
	for {
		fmt.Println(`Welcome to SatTV!

		1. View current balance in the account 
		2. Recharge Account \n
		3. View available packs, channels and services
		4. Subscribe to base packs
		5. Add channels to an existing subscription
		6. Subscribe to special services
		7. View current subscription details
		8. Update email and phone number for notifications
		9. Exit`)
		scanner.Scan()
		choice, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		switch choice {
		case 1:
			promptViewBalance()
		case 2:
			promptRecharge()
		case 3:
			viewAllServices()
		case 4:
			promptSubscribeBase()
		case 5:
			promptAddChannels()
		case 6:
			promptSubscribeSpecial()
		case 7:
			viewSubscription()
		case 8:
			promptUpdateInfo()
		case 9:
			break Loop
		default:
			fmt.Println("Invalid Choice. Please try again")
		}
	}
}
