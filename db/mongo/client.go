package mongo

import(
	"fmt"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/domain"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/db"
	"context"
    "time"
	"log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"regexp"
)
func init(){
	db.RegisterDataStore("mongo",NewClient)
}
func NewClient() (db.DataStore,error){
	ctx,cancel:=context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client1, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil{
		fmt.Println("The error is",err)
		return nil,err
	}
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client1.Ping(ctx, readpref.Primary())
	if err != nil{
		fmt.Println("The error is",err)
		return &client{dbc:nil},err
	}
	return &client{dbc:client1.Database("user_app")},nil
}

type client struct{
	dbc *mongo.Database
}

func (c *client) AddUser(user *domain.User) (string, error){
	collection := c.dbc.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_,err:=collection.InsertOne(ctx,bson.D{
		{Key:"_id",Value:user.ID},
		{Key:"name", Value:user.Name},
		{Key:"address", Value:user.Address}, 
		{Key:"created_at", Value:user.CreatedAt}, 
	})
	if err != nil{
		fmt.Println("The error is",err)
	}
	return user.ID,nil
}
func (c *client) ListUsers(filters map[string]interface{},limit int64) ([]*domain.User,error){
	var (
		cur *mongo.Cursor
		err error
	)
	options := options.Find().SetLimit(limit)
	cur, err = c.dbc.Collection("users").Find(context.Background(),applyFilter(filters),options)
    if err != nil {
        return nil, err
     }
	resultArray:=[]*domain.User{}
	for cur.Next(context.Background()) {
		var result1 *domain.User	
		err := cur.Decode(&result1)
		if err != nil { log.Fatal(err) }
		resultArray=append(resultArray,result1)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return resultArray,nil
}
func (c *client) DeleteUser(id string) error{
	c.dbc.Collection("users").DeleteOne(context.Background(), bson.M{"_id": id})
    return nil
}

func (c *client) ViewUser(id string) (*domain.User,error){
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var userInfo domain.User
	if err := c.dbc.Collection("users").FindOne(ctx, bson.M{"_id" : id}).Decode(&userInfo); err !=nil{
		return nil,&domain.Error{Code:404,Message:"User doesn't exist"}
	}
	return &userInfo,nil
}

func applyFilter(filterMap map[string]interface{}) map[string]interface{} {
    for k, v := range filterMap {
        switch mval := v.(type) {
        case []string:
            filterMap[k] = bson.M{"$in": mval}
        case string:
            // support searching by name using case-insensitive matching
            if k == "name" {
                filterMap[k] = bson.M{"$regex": primitive.Regex{Pattern: "^" + regexp.QuoteMeta(mval) + "$", Options: "i"}}
            }
        }
    }
    return filterMap
}