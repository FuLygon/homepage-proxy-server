package models

const (
	reqTypeContainer    = "GetDockerContainersSummary"
	reqTypeStack        = "GetStacksSummary"
	reqTypeBuild        = "GetBuildsSummary"
	reqTypeRepo         = "GetReposSummary"
	reqTypeAction       = "GetActionsSummary"
	reqTypeBuilder      = "GetBuildersSummary"
	reqTypeDeployment   = "GetDeploymentsSummary"
	reqTypeProcedure    = "GetProceduresSummary"
	reqTypeResourceSync = "GetResourceSyncsSummary"
)

// KomodoStatsResponse summarized response from Komodo API
type KomodoStatsResponse struct {
	Container    *KomodoContainerStats    `json:"container,omitempty"`
	Stack        *KomodoStackStats        `json:"stack,omitempty"`
	Build        *KomodoBuildStats        `json:"build,omitempty"`
	Repo         *KomodoRepoStats         `json:"repo,omitempty"`
	Action       *KomodoActionStats       `json:"action,omitempty"`
	Builder      *KomodoBuilderStats      `json:"builder,omitempty"`
	Deployment   *KomodoDeploymentStats   `json:"deployment,omitempty"`
	Procedure    *KomodoProcedureStats    `json:"procedure,omitempty"`
	ResourceSync *KomodoResourceSyncStats `json:"resource-sync,omitempty"`
}

type KomodoContainerStats struct {
	Total     int `json:"total"`
	Running   int `json:"running"`
	Stopped   int `json:"stopped"`
	Unhealthy int `json:"unhealthy"`
	Unknown   int `json:"unknown"`
}

func (m KomodoContainerStats) SummaryRequest() map[string]interface{} {
	return map[string]interface{}{
		"type":   reqTypeContainer,
		"params": struct{}{},
	}
}

type KomodoStackStats struct {
	Total     int `json:"total"`
	Running   int `json:"running"`
	Stopped   int `json:"stopped"`
	Down      int `json:"down"`
	Unhealthy int `json:"unhealthy"`
	Unknown   int `json:"unknown"`
}

func (m KomodoStackStats) SummaryRequest() map[string]interface{} {
	return map[string]interface{}{
		"type":   reqTypeStack,
		"params": struct{}{},
	}
}

type KomodoBuildStats struct {
	Total    int `json:"total"`
	Ok       int `json:"ok"`
	Failed   int `json:"failed"`
	Building int `json:"building"`
	Unknown  int `json:"unknown"`
}

func (m KomodoBuildStats) SummaryRequest() map[string]interface{} {
	return map[string]interface{}{
		"type":   reqTypeBuild,
		"params": struct{}{},
	}
}

type KomodoRepoStats struct {
	Total    int `json:"total"`
	Ok       int `json:"ok"`
	Cloning  int `json:"cloning"`
	Pulling  int `json:"pulling"`
	Building int `json:"building"`
	Failed   int `json:"failed"`
	Unknown  int `json:"unknown"`
}

func (m KomodoRepoStats) SummaryRequest() map[string]interface{} {
	return map[string]interface{}{
		"type":   reqTypeRepo,
		"params": struct{}{},
	}
}

type KomodoActionStats struct {
	Total   int `json:"total"`
	Ok      int `json:"ok"`
	Running int `json:"running"`
	Failed  int `json:"failed"`
	Unknown int `json:"unknown"`
}

func (m KomodoActionStats) SummaryRequest() map[string]interface{} {
	return map[string]interface{}{
		"type":   reqTypeAction,
		"params": struct{}{},
	}
}

type KomodoBuilderStats struct {
	Total int `json:"total"`
}

func (m KomodoBuilderStats) SummaryRequest() map[string]interface{} {
	return map[string]interface{}{
		"type":   reqTypeBuilder,
		"params": struct{}{},
	}
}

type KomodoDeploymentStats struct {
	Total       int `json:"total"`
	Running     int `json:"running"`
	Stopped     int `json:"stopped"`
	NotDeployed int `json:"not_deployed"`
	Unhealthy   int `json:"unhealthy"`
	Unknown     int `json:"unknown"`
}

func (m KomodoDeploymentStats) SummaryRequest() map[string]interface{} {
	return map[string]interface{}{
		"type":   reqTypeDeployment,
		"params": struct{}{},
	}
}

type KomodoProcedureStats struct {
	Total   int `json:"total"`
	Ok      int `json:"ok"`
	Running int `json:"running"`
	Failed  int `json:"failed"`
	Unknown int `json:"unknown"`
}

func (m KomodoProcedureStats) SummaryRequest() map[string]interface{} {
	return map[string]interface{}{
		"type":   reqTypeProcedure,
		"params": struct{}{},
	}
}

type KomodoResourceSyncStats struct {
	Total   int `json:"total"`
	Ok      int `json:"ok"`
	Syncing int `json:"syncing"`
	Pending int `json:"pending"`
	Failed  int `json:"failed"`
	Unknown int `json:"unknown"`
}

func (m KomodoResourceSyncStats) SummaryRequest() map[string]interface{} {
	return map[string]interface{}{
		"type":   reqTypeResourceSync,
		"params": struct{}{},
	}
}
