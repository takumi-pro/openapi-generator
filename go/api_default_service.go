/*
 * todos
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"errors"
	"net/http"

	"github.com/GIT_USER_ID/GIT_REPO_ID/infrastructure"
)

// DefaultApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService() DefaultApiServicer {
	return &DefaultApiService{}
}

// GetTask - Fetch All Task
func (s *DefaultApiService) GetTask(ctx context.Context) (ImplResponse, error) {
	// TODO - update GetTask with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetTask200Response{}) or use other options such as http.Ok ...
	//return Response(200, GetTask200Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetTask method not implemented. Let's implement GetTask !!")
}

// GetTaskTaskId - Fetch Task
func (s *DefaultApiService) GetTaskTaskId(ctx context.Context, taskId interface{}) (ImplResponse, error) {
	task := Task{}
	result := infrastructure.Db.WithContext(ctx).First(&task)

	if result.Error != nil {
		return Response(500, nil), errors.New("test")
	}

	return Response(200, Task{
		Id:          task.Id,
		Title:       task.Title,
		Description: task.Description,
	}), nil
}

// PostTask - Create New Task
func (s *DefaultApiService) PostTask(ctx context.Context, task Task) (ImplResponse, error) {
	// TODO - update PostTask with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("PostTask method not implemented")
}

// PutTaskTaskId - Update Task
func (s *DefaultApiService) PutTaskTaskId(ctx context.Context, taskId interface{}, task Task) (ImplResponse, error) {
	// TODO - update PutTaskTaskId with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("PutTaskTaskId method not implemented")
}
