package model

import "github.com/flipped-aurora/gva/global"

type WorkflowStartPoint struct {
	WorkflowEdgeID string
	global.Model
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Index int     `json:"index"`
}

type WorkflowEndPoint struct {
	WorkflowEdgeID string
	global.Model
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Index int     `json:"index"`
}
