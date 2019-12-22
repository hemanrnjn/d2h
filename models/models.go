package models

type Pack struct {
	Name  string
	Ch    []Channel
	Price int
}

type Channel struct {
	Ch    string
	Price int
}

type Service struct {
	Name  string
	Price int
}

type Subscription struct {
	Ch   []Channel
	Svc  []Service
	Pack string
}

type User struct {
	Balance     int
	Email       string
	Phonenumber string
	Subs        Subscription
}
