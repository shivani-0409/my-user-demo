package service
import(
	"fmt"
	"time"
	// "github.com/go-openapi/swag"	 
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/domain"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/db"
	"github.com/segmentio/ksuid"
	// "github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/gen/models"
	// "github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/gen/restapi/operations/users"
)

type manager struct{
	// ds map[int]*domain.User
	ds db.DataStore
}

func (m *manager) CreateUser(input *domain.User) (string,error){
	// n:=input.Name
	// add:=input.Address
	// identity:=input.ID
	// userInfo:= []*models.User{{Address:add,ID:identity,Name:n}}
	input.CreatedAt=time.Now().UTC()
	input.ID=ksuid.New().String()
	// for i:=0;i<=len(userInfo)-1;i++{
		//m.ds[input.ID]=input
		if len(input.Name) <=2 {
			return "",&domain.Error{Code:400,Message:"Name should be atleast 3 chars long"}
		}
		return m.ds.AddUser(input)
	// }
	//fmt.Println("The added user is",userInfo)
	//save this userInfo object into database
}

func (m *manager) ListUsers(queryName *string, queryLimit *int32) ([]*domain.User,error){
	//n can be any dynamic name
	// for i:=0;i<=len(m.ds)-1;i++{
		
	// }
	// var userInfo =[]*domain.User{}
	// for _,value := range m.ds{
	// 	userInfo = append(userInfo, value)
	// }
	// fmt.Println("The value of userInfo is",userInfo)
	// // n:="Shivani Sharma 1123"
	// // userInfo:= []*models.User{{Address:"MyAdd",ID:21,Name:&n}}
	// return userInfo
	//user,err:=m.ds.ListUsers(queryName,queryLimit)
	// if err !=nil{
	// 	return nil
	// } 
	return m.ds.ListUsers(queryName,queryLimit)
	//traverse the map, Range
}

// func (m *manager) DeleteUser(i int64) *models.Error{
// // func (m *manager) DeleteUser(i int64){
// 	//if i exists then del and ret nil else create a struct of models.err and send values in code and value
// 	for _,value:= range m.ds{
// 		if(value == m.ds[i]){
// 			delete(m.ds,i)
// 			return nil
// 		}
// 	  }
// 	  message:="User is not found"
// 	  return &models.Error{Code:404,Message:&message}
// }

func (m *manager) DeleteUser(i string) error{
	// if _, ok := m.ds[i]; ok{
	// 		  delete(m.ds,i)
	// 		   return nil
	// 	}
	// 	 return &domain.Error{Code:404,Message:"User doesn't exist"}
	_,err:=m.ds.ViewUser(i)
	if err != nil{
		return &domain.Error{Code:404,Message:"User doesn't exist"}
	}
	if err := m.ds.DeleteUser(fmt.Sprint(i)); err != nil{
		return &domain.Error{Code:404,Message:"User doesn't exist"}
	}
	return nil
  }

  func (m *manager) ViewUser(i string) (*domain.User,error){
	fmt.Printf("The type of id is %T",fmt.Sprint(i))
	// user,err := m.ds.ViewUser(i)
	// if err != nil{
	// 	return nil,&domain.Error{Code:404,Message:"User doesn't exist"}
	// }
	return m.ds.ViewUser(i)
  }
   
   