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
	// "go.mongodb.org/mongo-driver/bson/primitive"
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
	//Insert
	
	_,err:=collection.InsertOne(ctx,bson.D{
		{Key:"_id",Value:user.ID},
		{Key:"name", Value:user.Name},
		{Key:"address", Value:user.Address}, 
		{Key:"created_at", Value:user.CreatedAt}, 
	})
	if err != nil{
		fmt.Println("The error is",err)
	}

	//id := res.InsertedID
	//stringObjectID := id.(primitive.ObjectID).Hex()
	//fmt.Println("The ID is",id)
	return user.ID,nil
}
func (c *client) ListUsers(queryName *string,queryLimit *int32) ([]*domain.User,error){
	//Find
	collection := c.dbc.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var (
		cur *mongo.Cursor
		err error
	)
	cur, err = collection.Find(ctx,bson.D{})
	
	if queryLimit !=nil{
		if *queryLimit != 0{
			if queryName != nil{
				if *queryName != ""{
					options := options.Find()
					options.SetLimit(int64(*queryLimit))
					// cur, err = collection.Find(ctx,bson.D{},options)
					cur, err = collection.Find(ctx,bson.M{"name" : *queryName},options)
					fmt.Println("The name query is",*queryName)
				}
			} else{
					options := options.Find()
					options.SetLimit(int64(*queryLimit))
					cur, err = collection.Find(ctx,bson.D{},options)
			}
		}
	}
	if err != nil { log.Fatal(err) }
	defer cur.Close(ctx)
	resultArray:=[]*domain.User{}
	for cur.Next(ctx) {
		var result bson.D
		var result1 *domain.User	
		err := cur.Decode(&result1)
		if err != nil { log.Fatal(err) }
		resultArray=append(resultArray,result1)
		fmt.Println("The result is",result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("The final result(List User) is",resultArray)
	return resultArray,nil
}
func (c *client) DeleteUser(id string) error{
	//Delete 
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	
	// cur, err := c.dbc.Collection("users").Find(ctx, bson.M{"_id" : id})
	// if err != nil { log.Fatal(err) }
	// defer cur.Close(ctx)
	// resultArray:=[]*domain.User{}
	// for cur.Next(ctx) {
	// 	// var result bson.D
	// 	var result1 *domain.User	
	// 	err := cur.Decode(&result1)
	// 	if err != nil { log.Fatal(err) }
	// 	resultArray=append(resultArray,result1)
	// }
	// if len(resultArray) == 0{
	// 	return &domain.Error{Code:404,Message:"User doesn't exist"}
	// }
	c.dbc.Collection("users").DeleteOne(context.Background(), bson.M{"_id": id})
    return nil
}

func (c *client) ViewUser(id string) (*domain.User,error){
	//View User 
	// collection := c.dbc.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	//objID, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	return nil,err
	//   }
	var userInfo domain.User
	if err := c.dbc.Collection("users").FindOne(ctx, bson.M{"_id" : id}).Decode(&userInfo); err !=nil{
		return nil,&domain.Error{Code:404,Message:"User doesn't exist"}
	}
	
	// if err != nil { log.Fatal(err) }
	// defer cur.Close(ctx)
	// resultArray:=*domain.User
	// for cur.Next(ctx) {
	// 	var result bson.D
	// 	var result1 *domain.User	
	// 	err := cur.Decode(&result1)
	// 	if err != nil { log.Fatal(err) }
	// 	resultArray=append(resultArray,result1)
	// 	fmt.Println("The viewed result is",result)
	// }
	// if err := cur.Err(); err != nil {
	// 	log.Fatal(err)
	// }
	return &userInfo,nil
}