package plan

type Action struct {
	Type     string `json:"type"`
	Identity string `json:"identity"`
}

type Plan struct {
	Actions []Action `json:"actions"`
}
