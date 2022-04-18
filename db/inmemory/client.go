package inmemory
import(
	"fmt"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/domain"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/db"
	"strings"
)

func init(){
	db.RegisterDataStore("inmemory",NewClient)
}

func NewClient() (db.DataStore,error){
	return &client{ds:make(map[string]*domain.User)},nil
}

type client struct{
	ds map[string]*domain.User
}

func (c *client) AddUser(user *domain.User) (string, error){
	c.ds[user.ID]=user
	return user.ID,nil
}
func (c *client) ListUsers(filters map[string]interface{}, limit int64) ([]*domain.User,error){
	var userInfo =[]*domain.User{}
	var lim int64 = 1
	for _,value := range c.ds{
		if filters["name"] !=nil{
			if strings.ToUpper(fmt.Sprint(filters["name"])) == strings.ToUpper(value.Name){
				userInfo = append(userInfo, value)
			}
		} else{
			userInfo = append(userInfo, value)
		}
		if limit != 0{
			if lim == limit{
				break
			}
			lim++
		}  
	}
	return userInfo,nil
}
func (c *client) DeleteUser(id string) error{
	if _, ok := c.ds[id]; ok{
		delete(c.ds,id)
		return nil
  }
return &domain.Error{Code:404,Message:"User doesn't exist"}
}

func (c *client) ViewUser(id string) (*domain.User,error){
	var userInfo *domain.User
	var ok bool
	if userInfo, ok = c.ds[id]; !ok{
		 return nil,&domain.Error{Code:404,Message:"User doesn't exist"}
  	}
	return userInfo,nil
}

