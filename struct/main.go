package main
import "fmt"
type Person struct{
	firstName string
	lastName string
	age int 
}
type Contact struct{
	Email string
	Phone string
}
type Address struct{
	HouseNumber int
	Area string
	state string
}
type Employee struct{
	person_Details Person
	person_Contact Contact
	person_Address Address
}
func main(){
    var person1 Person
	fmt.Println(person1)
	person1.firstName="person1"
	person1.lastName="person1"
	person1.age=23
	fmt.Println(person1)
    
	person2:=Person{
		firstName:"person2",
		lastName:"person2",
		age:24,
	}
	fmt.Println(person2)

	var person3=new(Person)
	person3.firstName="person3"
	person3.lastName="person3"
	person3.age=25
	fmt.Println(person3)
	fmt.Println(person3.firstName)

	 var employee1 Employee
	 employee1.person_Details=person2
	 employee1.person_Contact=Contact{
		Email: "contact@gmail.com",
		Phone: "1234567890",
	 }
	 employee1.person_Address.state="state"
	 employee1.person_Address.Area="area"
	 employee1.person_Address.HouseNumber=103
	 fmt.Println(employee1)
	 fmt.Println(employee1.person_Details)
	 fmt.Println(employee1.person_Contact)
	 fmt.Println(employee1.person_Address)
}