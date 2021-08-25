package services

import (
	"WebDevelopmentTodd/Webservices/REST_MS_GO/golang-microservices/src/api/config"
	"WebDevelopmentTodd/Webservices/REST_MS_GO/golang-microservices/src/api/domain/github"
	"WebDevelopmentTodd/Webservices/REST_MS_GO/golang-microservices/src/api/domain/repositories"
	"WebDevelopmentTodd/Webservices/REST_MS_GO/golang-microservices/src/api/providers/github_provider"
	"WebDevelopmentTodd/Webservices/REST_MS_GO/golang-microservices/src/api/utils/errors"
	"fmt"
	"net/http"
	"sync"
)

type repoService struct{}

type repoServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
}

var (
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

func (s *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}
	result := repositories.CreateRepoResponse{
		Id:    response.Id,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil
}

func (s *repoService) CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError) {
	fmt.Println("Inside repository service's create repos")
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	defer close(output)
	var wg sync.WaitGroup

	go s.handleRepoResult(&wg, input, output)
	for _, r := range request {
		wg.Add(1)
		go s.CreateRepoConcurrent(r, input)
	}
	wg.Wait()
	close(input)
	result := <-output
	successCreations := 0

	for _, res := range result.Results {
		if res.Error == nil {
			successCreations++
		}
	}
	if successCreations == 0 {
		result.Status = result.Results[0].Error.Status()
	} else if successCreations == len(request) {
		result.Status = http.StatusCreated
	} else {
		result.Status = http.StatusPartialContent
	}

	return result, nil
}

func (s *repoService) handleRepoResult(wg *sync.WaitGroup, input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse) {
	fmt.Println("Inside Handle Repo Result")
	var results repositories.CreateReposResponse
	for incomingEvent := range input {
		repoResult := repositories.CreateRepositoriesResult{
			Response: incomingEvent.Response,
			Error:    incomingEvent.Error,
		}
		results.Results = append(results.Results, repoResult)
		wg.Done()
	}
	output <- results
}

func (s *repoService) CreateRepoConcurrent(input repositories.CreateRepoRequest, output chan repositories.CreateRepositoriesResult) {
	fmt.Println("Inside CreateRepoConcurrent")
	if err := input.Validate(); err != nil {
		fmt.Println("Validation Failed inside CreateRepoConcurrent")
		output <- repositories.CreateRepositoriesResult{
			Error: err,
		}
		return
	}
	result, err := s.CreateRepo(input)
	if err != nil {
		fmt.Println("Create Repo Failed inside CreateRepoConcurrent")
		output <- repositories.CreateRepositoriesResult{
			Error: err,
		}
		return
	}
	output <- repositories.CreateRepositoriesResult{Response: result}

}
