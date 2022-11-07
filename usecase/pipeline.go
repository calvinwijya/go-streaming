package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/calvinwijya/go-streaming/model"
	"github.com/calvinwijya/go-streaming/repository"
)

type pipelineUsecase struct {
	repository repository.Repository
}

type Param struct {
	Repository repository.Repository
}

func NewPipelineUsecase(param Param) (Usecase, error) {
	if param.Repository == nil {
		err := fmt.Errorf("invalid params are provided")
		return nil, err
	}

	pipelineUsecase := &pipelineUsecase{
		repository: param.Repository,
	}

	return pipelineUsecase, nil
}

func (p pipelineUsecase) ProcessData(events []interface{}) {
	//Print length of events received
	fmt.Println("Received events count", len(events))

	// Make data model slice of events length as capacity
	modelInfo := make([]model.DataModel, 0, len(events))

	// Loop through input events, unmarshal Json to looping model and append to model slice
	for _, event := range events {
		var dataModel model.DataModel

		err := json.Unmarshal(event.([]byte), &dataModel)
		if err != nil {
			fmt.Println("failed to unmarshal events to dataModel")
			continue
		}

		modelInfo = append(modelInfo, dataModel)

	}

	// Call function processDataToMap(), to process data to map
	data := p.processDataToMap(modelInfo)

	// Call Repository Store method with input map data
	fmt.Println("trigger the repository to store data")
	err := p.repository.Store(data)
	if err != nil {
		fmt.Println("error processing Store() method", err)
	}
}

func (p pipelineUsecase) processDataToMap(modelInfo []model.DataModel) []map[string]interface{} {
	data := []map[string]interface{}{}

	for _, model := range modelInfo {
		record := map[string]interface{}{}
		record["pid"] = model.Pid
		record["recommended_pids"] = model.RecommendedPids
		data = append(data, record)
	}

	return data
}
