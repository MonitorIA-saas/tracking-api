package track

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Process struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	OcorrenciaID   primitive.ObjectID `json:"ocorrencia_id,omitempty" bson:"ocorrencia_id,omitempty"`
	PID            int                `json:"PID" bson:"PID"`
	Priority       int                `json:"priority" bson:"priority"`
	AlocatedMemory float64            `json:"alocated_memory" bson:"alocated_memory"`
	State          int                `json:"state" bson:"state"`
	CPUUsage       float64            `json:"cpu_usage" bson:"cpu_usage"`
	GPUUsage       float64            `json:"gpu_usage" bson:"gpu_usage"`
	IO             []string           `json:"io" bson:"io"`
	CPUTime        int                `json:"cpu_time" bson:"cpu_time"`
	CPURegisters   int                `json:"cpu_registers" bson:"cpu_registers"`
	ProgramCounter int                `json:"program_counter" bson:"program_counter"`
	Timestamp      time.Time          `json:"timestamp" bson:"timestamp"`
	ExecutionTime  time.Time          `json:"execution_time" bson:"execution_time"`
}

type Occurency struct {
	ID                 primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CollaboratorPSQLID int                `json:"collaborator_psql_id" bson:"collaborator_psql_id"`
	URL                string             `json:"url" bson:"url"`
	InfractionIndex    float64            `json:"infraction_index" bson:"infraction_index"`
	Timestamp          time.Time          `json:"timestamp" bson:"timestamp"`
	Status             string             `json:"status" bson:"status"`
}
