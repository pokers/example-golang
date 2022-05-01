package receiver

import (
	"fmt"
)

type stUser struct{
	id uint32
	name string
	age uint8
}
type User struct{
	Id uint32
	Name string
	Age uint8
}

func (user stUser) getUserInto() string {
	return fmt.Sprintf("ID[%d] Name[%s] Age[%d]", user.id, user.name, user.age)
}
func (user *stUser) increaseAge() {
	fmt.Printf("AGE : %d\n", user.age)
	user.age++
	fmt.Printf("After AGE : %d\n", user.age)
}
func (user stUser) String() string {
	return fmt.Sprintf("ID[%d] Name[%s] Age[%d]", user.id, user.name, user.age)
}


func (user User) getUserInto() string {
	return fmt.Sprintf("ID[%d] Name[%s] Age[%d]", user.Id, user.Name, user.Age)
}

func GetUser() string {
	var user stUser = stUser{
		id: 0,
		name: "Seo JoongSu",
		age: 20,
	}
	user.increaseAge()
	return user.getUserInto()
}

func GetUsers() {
	var users []*stUser = []*stUser{}
	// users := []*stUser{}
	users = append(users, new(stUser))

	userA := new(stUser)
	userA.id=1
	userA.name = "JoongSu"
	userA.age = 22

	users = append(users, userA)
	
	
	fmt.Println(users)
	for index,user := range users {
		fmt.Printf("> user[%d] : %s\n", index, user)
	}
}

func GetUsers2() {
	var users [10]stUser = [10]stUser{
		{
			id: 0,
			name: "Seo JoongSu",
			age: 20,
		},
		{
			id: 0,
			name: "Seo JoongSu",
			age: 20,
		},
	}
	// users := []*stUser{}
	// users[0] = new(stUser)

	// userA := new(stUser)
	// userA.id=1
	// userA.name = "JoongSu"
	// userA.age = 22
	// users[3] = userA

	
	
	fmt.Println(users)
	for index := range(users) {
		var user stUser = users[index]
		fmt.Println("user[", index, "] ", user)
	}
}
