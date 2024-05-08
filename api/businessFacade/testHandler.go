package businessFacade

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/GeldNetworkMVP/GeldMVPBackend/dao"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
	log "github.com/sirupsen/logrus"
)

var dt = time.Now()

func SaveWorkflow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	dt := time.Now()
	var Workflow model.PlotChainData
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&Workflow)
	if err != nil {
		log.Println(err)
	}
	if Workflow.PlotID != "" {
		WorkflowObj = model.PlotChainData{
			PlotID: Workflow.PlotID,
			Status: Workflow.Status,
			Workflow: Workflow.Workflow{
				WorkflowID: Workflow.Workflow.WorkflowID,
				Status:     Workflow.Workflow.Status,
				Stages: Workflow.Workflow.Stages{
					StageID:   Workflow.Workflow.Stages.StageID,
					StageName: Workflow.Workflow.Stages.StageName,
					StageFieldNames: Workflow.Workflow.Stages.StageFieldNames{
						Field1: Workflow.Workflow.Stages.StageFieldNames.Field1,
						Field2: Workflow.Workflow.Stages.StageFieldNames.Field2,
						Field3: Workflow.Workflow.Stages.StageFieldNames.Field3,
					},
				},
			},
		}

		WorkflowResponse := model.WorkflowResponse{
			WorkflowID: Workflow.Workflow.WorkflowID,
			Status:     "Successfully Created",
		}
		object := dao.Connection{}
		err1, err2 := object.SaveWorkflow(WorkflowObj)
		if err1 != nil && err2 != nil {
			log.Error("NFT not inserted : ", err1, err2)
		}
		if err1 == nil && err2 != nil {
			log.Error("NFT not inserted into StellarNFT Collection : ", err2)
		}
		if err1 != nil && err2 == nil {
			log.Error("NFT not inserted into Marketplace Collection : ", err1)
		} else {
			log.Error("NFT inserted to the collection")
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(WorkflowResponse)
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		response := model.Error{Message: "Data Template Integrity Issue: PlotID not defined"}
		json.NewEncoder(w).Encode(response)
	}
}
