package repositories

import (
	"Submission_Service/dtos"
	"Submission_Service/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SubmissionRepository interface {
	CreateSubmission(submission *dtos.CreateSubmissionRequest) (dtos.SubmissionResponse , error)
	GetSubmissionByID(id string) (*dtos.SubmissionResponse, error)
	UpdateSubmission(id string, submission *dtos.CreateSubmissionRequest) (*dtos.SubmissionResponse, error)
	DeleteSubmission(id string) error
}

type SubmissionRepositoryImpl struct {
	collection *mongo.Collection
}

func NewSubmissionRepository(_client *mongo.Client) SubmissionRepository {
	return &SubmissionRepositoryImpl{
		collection: _client.Database("Leetcode").Collection("submissions"),
	}
}


func (r *SubmissionRepositoryImpl) CreateSubmission(submission *dtos.CreateSubmissionRequest) (dtos.SubmissionResponse , error) {
	submisssionModel := &models.Submission{
		ProblemID : &submission.ProblemId,
		Code: &submission.Code,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	result , err := r.collection.InsertOne(context.TODO() , submisssionModel)
	if err != nil {
		fmt.Println("Error inserting submission:", err)
		return dtos.SubmissionResponse{} , err
	}
	return dtos.SubmissionResponse{
		Id: result.InsertedID.(string),
		Status: "Created",
		ProblemId: submission.ProblemId,
		Code: submission.Code,
		CreatedAt: submisssionModel.CreatedAt.String(),
		UpdatedAt: submisssionModel.UpdatedAt.String(),
	}, nil
}

func (r *SubmissionRepositoryImpl) GetSubmissionByID(id string) (*dtos.SubmissionResponse, error) {
	// Implementation for retrieving a submission by ID from the database
	var submissionModel models.Submission
	err := r.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&submissionModel)
	if err != nil {
		fmt.Println("Error fetching submission:", err)
		return nil, err
	}
	return &dtos.SubmissionResponse{
		Id:        *submissionModel.ID,
		ProblemId: *submissionModel.ProblemID,
		Code:     *submissionModel.Code,
		Status:   *submissionModel.Status,
		CreatedAt: submissionModel.CreatedAt.String(),
		UpdatedAt: submissionModel.UpdatedAt.String(),
	}, nil
}

func (r *SubmissionRepositoryImpl) UpdateSubmission(id string, submission *dtos.CreateSubmissionRequest) (*dtos.SubmissionResponse, error) {
    updateFields := bson.M{}
    
    if submission.ProblemId != "" {
        updateFields["problem_id"] = submission.ProblemId
    }
    if submission.Code != "" {
        updateFields["code"] = submission.Code
    }

    updateFields["updated_at"] = time.Now()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectID}
    update := bson.M{"$set": updateFields}

    result , err := r.collection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
		fmt.Println("Error updating submission:", err)
        return nil, err
    }

    var updated models.Submission
    err = r.collection.FindOne(context.TODO(), filter).Decode(&updated)
    if err != nil {
        return nil, err
    }

    return &dtos.SubmissionResponse{
        Id:        result.UpsertedID.(string),
        ProblemId: *updated.ProblemID,
        Code:      *updated.Code,
        Status:    *updated.Status,
        CreatedAt: result.UpsertedID.(time.Time).Format(time.RFC3339),
        UpdatedAt: result.UpsertedID.(time.Time).Format(time.RFC3339),
    }, nil
}


func (r *SubmissionRepositoryImpl) DeleteSubmission(id string) error {
	_, err := r.collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		fmt.Println("Error deleting submission:", err)
		return err
	}
	return nil
}