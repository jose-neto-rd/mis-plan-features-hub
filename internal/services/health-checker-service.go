package services

type HealthCheckerService struct {
	InstanceId string
}

func (h *HealthCheckerService) CheckLiveness() bool {
	return true
}

func (h *HealthCheckerService) CheckReadiness() map[string]interface{} {
	dependencies := map[string]string{
		"pubsub": "ok",
	}

	/**
	TODO CHECK HEALTH PUB/SUB
	if err := pubsub.CheckPubSubConnection(); err != nil {
		dependencies["pubsub"] = "error"
	}
	*/

	status := "ok"
	for _, depStatus := range dependencies {
		if depStatus == "error" {
			status = "error"
			break
		}
	}

	instanceID := h.InstanceId
	if instanceID == "" {
		instanceID = "local"
	}

	return map[string]interface{}{
		"status":       status,
		"instance_id":  instanceID,
		"dependencies": dependencies,
	}
}
