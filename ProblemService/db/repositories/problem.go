package repositories

import (
	"context"
	"fmt"
	"leetcode/dtos"
	"leetcode/models"
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
		Id :          result.InsertedID.(primitive.ObjectID).Hex(),
		Title:      problem.Title,
		Description: problem.Description,
		Editorial:   problem.Editorial,
		CreatedAt:   result.InsertedID.(primitive.ObjectID).Timestamp().String(),
		UpdatedAt:   result.InsertedID.(primitive.ObjectID).Timestamp().String(),
	}, nil
}

func (r *ProblemRepositoryImpl) GetProblemById(id string) (*dtos.ProblemResponse, error) {
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
	return problem, nil
}

func (r *ProblemRepositoryImpl) UpdateProblem(id string, problem *models.Problem) (*dtos.ProblemResponse, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("Error updating problem:", err)
		return nil, err
	}

	updatedResult, err := r.collection.UpdateOne(context.TODO(), bson.M{"_id": objId}, bson.M{"$set": problem})
	if err != nil {
		fmt.Println("Error updating problem:", err)
		return nil, err
	}


	fmt.Println("Problem updated successfully")
	return &dtos.ProblemResponse{
		Id:         updatedResult.UpsertedID.(primitive.ObjectID).Hex(),
		Title:      problem.Title,
		Description: problem.Description,
		Editorial:   problem.Editorial,
		CreatedAt:   updatedResult.UpsertedID.(primitive.ObjectID).Timestamp().String(),
		UpdatedAt:   updatedResult.UpsertedID.(primitive.ObjectID).Timestamp().String(),
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
