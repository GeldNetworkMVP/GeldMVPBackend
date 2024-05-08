package businessFacade

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/GeldNetworkMVP/GeldMVPBackend/dao/dataTemplates"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
	log "github.com/sirupsen/logrus"
)

var dt = time.Now()

func SaveWorkflow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var Workflow model.PlotChainData
	var WorkflowObj model.PlotChainData
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
			Workflow: model.Workflow{
				WorkflowID: Workflow.Workflow.WorkflowID,
				Status:     Workflow.Workflow.Status,
				Stages: model.Stages{
					StageID:   Workflow.Workflow.Stages.StageID,
					StageName: Workflow.Workflow.Stages.StageName,
					StageFieldNames: model.StageFieldNames{
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
		object := dataTemplates.Connection{}
		err1 := object.SaveWorkflow(WorkflowObj)
		if err1 != nil {
			log.Error("Workflow not inserted : ", err1)
		} else {
			log.Error("Workflow inserted to the collection")
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
