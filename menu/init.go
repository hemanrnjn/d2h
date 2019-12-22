package menu

import "github.com/hemanrnjn/d2h/models"

func init() {
	User = models.User{
		Balance:     100,
		Email:       "",
		Phonenumber: "",
		Subs:        models.Subscription{},
	}

	AllPacks = []models.Pack{
		{
			Name: "Silver Pack",
			Ch: []models.Channel{
				{
					Ch:    "Zee",
					Price: 10,
				},
				{
					Ch:    "Sony",
					Price: 15,
				},
				{
					Ch:    "Star Plus",
					Price: 20,
				},
			},
			Price: 50,
		},
		{
			Name: "Gold Pack",
			Ch: []models.Channel{
				{
					Ch:    "Zee",
					Price: 10,
				},
				{
					Ch:    "Sony",
					Price: 15,
				},
				{
					Ch:    "Star Plus",
					Price: 20,
				},
				{
					Ch:    "Discovery",
					Price: 10,
				},
				{
					Ch:    "Nat Geo",
					Price: 20,
				},
			},
			Price: 100,
		},
	}

	AllChannels = []models.Channel{
		{
			Ch:    "Zee",
			Price: 10,
		},
		{
			Ch:    "Sony",
			Price: 15,
		},
		{
			Ch:    "Star Plus",
			Price: 20,
		},
		{
			Ch:    "Discovery",
			Price: 10,
		},
		{
			Ch:    "Nat Geo",
			Price: 20,
		},
	}

	AllServices = []models.Service{
		{
			Name:  "LearnEnglish Service",
			Price: 200,
		},
		{
			Name:  "LearnCooking Service",
			Price: 100,
		},
	}

}
