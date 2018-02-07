package client

const (
	AlertConditionType                    = "alertCondition"
	AlertConditionFieldLastTransitionTime = "lastTransitionTime"
	AlertConditionFieldLastUpdateTime     = "lastUpdateTime"
	AlertConditionFieldMessage            = "message"
	AlertConditionFieldReason             = "reason"
	AlertConditionFieldStatus             = "status"
	AlertConditionFieldType               = "type"
)

type AlertCondition struct {
	LastTransitionTime string `json:"lastTransitionTime,omitempty"`
	LastUpdateTime     string `json:"lastUpdateTime,omitempty"`
	Message            string `json:"message,omitempty"`
	Reason             string `json:"reason,omitempty"`
	Status             string `json:"status,omitempty"`
	Type               string `json:"type,omitempty"`
}
