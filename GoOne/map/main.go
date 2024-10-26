package main
import "fmt"
func main(){
  studentGrades:=make(map[string]int)
  studentGrades["user1"]=25
  studentGrades["user2"]=26
  studentGrades["user3"]=27
  studentGrades["user4"]=28

  fmt.Println("Marks of user3:" ,studentGrades["user3"])
  studentGrades["user2"]=100
  fmt.Println("Marks of user2:",studentGrades["user2"])
  delete(studentGrades,"user1")
  fmt.Println("Marks of user1:",studentGrades["user1"])
  score,isPresent:=studentGrades["user5"]
  fmt.Println("user 5 Score:",score)
  fmt.Println("is user 5 exist:",isPresent)

  Score,IsPresent:=studentGrades["user4"]
  fmt.Println("Grade of user4:",Score)
  fmt.Println("is User4 Prsent:",IsPresent)

  for key,value:=range studentGrades{
	fmt.Printf("name is %s, marks is %d\n",key,value)
  }
  person:=map[string]int{"person1":12,"person2":13,"person":14}
  
  for key,value:= range person{
	fmt.Println(key,value)
  }

}