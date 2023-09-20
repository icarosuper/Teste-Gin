package entities

type User struct {
	Name string
	//Birth time.Time
}

func NewUser(name string) User {
	return User{
		name,
		//utils.GetDate(21, time.Month(6), 2000),
	}
}
