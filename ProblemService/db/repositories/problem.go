package repositories

import (
	"context"
	"fmt"
	"leetcode/dtos"
	"leetcode/models"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProblemRepository interface {
	CreateProblem(problem *models.Problem) (*dtos.ProblemResponse, error)
	GetProblemById(id string) (*dtos.ProblemResponse, error)
	UpdateProblem(id string , problem *models.Problem) (*dtos.ProblemResponse, error)
	DeleteProblem(id string) error
	GetAllProblem() ([]*dtos.ProblemResponse, error)
	SearchProblem(query string) ([]*dtos.ProblemResponse, error)
}

type ProblemRepositoryImpl struct {
	collection *mongo.Collection
}

func NewProblemRepository(_client *mongo.Client) ProblemRepository {
	return &ProblemRepositoryImpl{
		collection: _client.Database("Problem").Collection("problems"),
	}
}

func (r *ProblemRepositoryImpl) CreateProblem(problem *models.Problem) (*dtos.ProblemResponse, error) {
	result, err := r.collection.InsertOne(context.TODO(), problem)

	if err != nil {
		fmt.Println("Error inserting problem:", err)
		return nil, err
	}


	return &dtos.ProblemResponse{
		Id:          result.InsertedID.(primitive.ObjectID).Hex(),
		Title:      *problem.Title,
		Description: *problem.Description,
		Editorial:   *problem.Editorial,
		CreatedAt:   result.InsertedID.(primitive.ObjectID).Timestamp().String(),
		UpdatedAt:   result.InsertedID.(primitive.ObjectID).Timestamp().String(),
	}, nil
}

func (r *ProblemRepositoryImpl) GetProblemById(id string) (*dtos.ProblemResponse, error) {
	fmt.Println("Fetching problem with id:", id)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("Error converting id to ObjectID:", err)
		return nil, err
	}

	problem := &dtos.ProblemResponse{}
	err = r.collection.FindOne(context.TODO(), bson.M{"_id": objId}).Decode(problem)

	if err != nil {
		fmt.Println("Error finding problem:", err)
		return nil, err
	}
	return problem , nil
}

func (r *ProblemRepositoryImpl) UpdateProblem(id string, problem *models.Problem) (*dtos.ProblemResponse, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("Error updating problem:", err)
		return nil, err
	}

	update := bson.M{}
    val := reflect.ValueOf(problem).Elem()
    typ := val.Type()

    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)
        if !field.IsNil() {
            bsonTag := typ.Field(i).Tag.Get("bson")
            if bsonTag != "" {
                update[bsonTag] = field.Interface()
            }
        }
    }
	update["updated_at"] = time.Now().String()
	_, err = r.collection.UpdateOne(context.TODO(), bson.M{"_id": objId}, bson.M{"$set": update})
	if err != nil {
		fmt.Println("Error updating problem:", err)
		return nil, err
	}

	fmt.Println("Problem updated successfully")
	 return &dtos.ProblemResponse{
		Id:          objId.Hex(), //TODO : Add proper respose structure , created and updated at
		Title:      "Title Updated",
		Description: "Description Updated",
		Editorial:   "Editorial Updated",
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}, nil
}

func (r *ProblemRepositoryImpl) DeleteProblem(id string) error {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		fmt.Println("Error converting id to ObjectID:", err)
		return err
	}

	_, err = r.collection.DeleteOne(context.TODO(), bson.M{"_id": objId})

	if err != nil {
		fmt.Println("Error deleting problem:", err)
		return err
	}
	fmt.Println("Problem deleted successfully")
	return nil
}

func (r *ProblemRepositoryImpl) GetAllProblem() ([]*dtos.ProblemResponse, error) {
	curs, err := r.collection.Find(context.TODO(), bson.M{})

	if err != nil {
		fmt.Println("Error in retrieving all problems !" , err)
		return nil , err
	}

	var problems []*dtos.ProblemResponse
	for curs.Next(context.TODO()) {
		var problem dtos.ProblemResponse
		if err := curs.Decode(&problem); err != nil {
			fmt.Println("Error decoding problem:", err)
			continue
		}
		problems = append(problems, &problem)
	}

	curs.Close(context.TODO())
	return problems, nil
}

func (r *ProblemRepositoryImpl) SearchProblem(query string) ([]*dtos.ProblemResponse, error) {
	filter := bson.D{
		{Key: "$text", Value: bson.D{
			{Key: "$search", Value: query},
		}},
	}
	//TODO --> Not Working , Creating Indexes left
	_, err := r.collection.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys: bson.D{
			{Key: "title", Value: "text"},
			{Key: "description", Value: "text"},
			{Key: "editorial", Value: "text"},
		},
	})

	if err != nil {
		fmt.Println("Error creating text index:", err)
		return nil, err
	}

	curs, err := r.collection.Find(context.TODO(), filter)

	if err != nil {
		fmt.Println("Error in retrieving all problems !", err)
		return nil, err
	}

	var problems []*dtos.ProblemResponse
	for curs.Next(context.TODO()) {
		var problem dtos.ProblemResponse
		if err := curs.Decode(&problem); err != nil {
			fmt.Println("Error decoding problem:", err)
			continue
		}
		problems = append(problems, &problem)
	}

	curs.Close(context.TODO())
	return problems, nil
}
