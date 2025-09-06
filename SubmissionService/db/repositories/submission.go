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
	CreateSubmission(submission *dtos.CreateSubmissionRequest) (dtos.SubmissionResponse, error)
	GetSubmissionByID(id string) (*dtos.SubmissionResponse, error)
	UpdateSubmission(id string, submission *dtos.CreateSubmissionRequest) (*dtos.SubmissionResponse, error)
	DeleteSubmission(id string) error
}

type SubmissionRepositoryImpl struct {
	collection *mongo.Collection
}

func NewSubmissionRepository(_collection *mongo.Collection) SubmissionRepository {
	return &SubmissionRepositoryImpl{
		collection: _collection,
	}
}


func (r *SubmissionRepositoryImpl) CreateSubmission(submission *dtos.CreateSubmissionRequest) (dtos.SubmissionResponse, error) {
	submissionModel := &models.Submission{
		ProblemID: submission.ProblemId,
		Code:      submission.Code,
		Language: submission.Language,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := r.collection.InsertOne(context.TODO(), submissionModel)
	if err != nil {
		fmt.Println("Error inserting submission:", err)
		return dtos.SubmissionResponse{}, err
	}

	// Convert ObjectID to string properly
	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return dtos.SubmissionResponse{}, fmt.Errorf("failed to convert InsertedID to ObjectID")
	}
	id := objectID.Hex()

	// Use value types instead of pointers for basic types
	createdAt := submissionModel.CreatedAt.Format(time.RFC3339)
	updatedAt := submissionModel.UpdatedAt.Format(time.RFC3339)

	return dtos.SubmissionResponse{
		Id:        id, // Use string instead of *string
		ProblemId: *submission.ProblemId,
		Code:      *submission.Code,
		Language: *submission.Language,
		Status:    dtos.Status("Pending"), 
		CreatedAt: createdAt,              
		UpdatedAt: updatedAt,             
	}, nil
}

func (r *SubmissionRepositoryImpl) GetSubmissionByID(id string) (*dtos.SubmissionResponse, error) {
	var submissionModel models.Submission
	err := r.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&submissionModel)
	if err != nil {
		fmt.Println("Error fetching submission:", err)
		return nil, err
	}

	CreatedAt := submissionModel.CreatedAt.String()
	UpdatedAt := submissionModel.UpdatedAt.String()
	return &dtos.SubmissionResponse{
		Id:        *submissionModel.ID,
		ProblemId: *submissionModel.ProblemID,
		Code:      *submissionModel.Code,
		Status:    dtos.Status(*submissionModel.Status),
		CreatedAt: CreatedAt,
		UpdatedAt: UpdatedAt,
	}, nil
}

func (r *SubmissionRepositoryImpl) UpdateSubmission(id string, submission *dtos.CreateSubmissionRequest) (*dtos.SubmissionResponse, error) {
	updateFields := bson.M{}

	if submission.ProblemId != nil {
		updateFields["problem_id"] = *submission.ProblemId
	}
	if submission.Code != nil {
		updateFields["code"] = *submission.Code
	}

	updateFields["updated_at"] = time.Now()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": updateFields}

	result, err := r.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println("Error updating submission:", err)
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("no submission found with id %s", id)
	}

	var updated models.Submission
	err = r.collection.FindOne(context.TODO(), filter).Decode(&updated)
	if err != nil {
		return nil, err
	}

	updateFields["created_at"] = updated.CreatedAt

	return &dtos.SubmissionResponse{
		Id:        id,
		ProblemId: *updated.ProblemID,
		Code:      *updated.Code,
		Status:    dtos.Status(*updated.Status),
		CreatedAt: updateFields["created_at"].(time.Time).String(),
		UpdatedAt: updateFields["updated_at"].(time.Time).String(),
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
