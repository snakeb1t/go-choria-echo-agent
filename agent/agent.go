package agent

import (
	"context"
	"github.com/choria-io/go-choria/choria"
	"github.com/choria-io/go-choria/plugin"
	"github.com/choria-io/go-choria/server"
	"github.com/choria-io/go-choria/server/agents"
	"github.com/choria-io/mcorpc-agent-provider/mcorpc"
)

type EchoRequest struct {
	Message string `json:"message" validate:"shellsafe"`
}

type EchoReply struct {
	Message string `json:"message"`
	TimeStamp int `json:"timestamp"`
}

var metadata = &agents.Metadata{
	Name:        "echo",
	Description: "Choria Echo Agent",
	Author:      "R.I.Pienaar <rip@devco.net>",
	Version:     "1.0.0",
	License:     "Apache-2",
	Timeout:     2,
	URL:         "http://choria.io",
}

func New(mgr server.AgentManager) (agents.Agent, error) {
	agent := mcorpc.New("echo", metadata, mgr.Choria(), mgr.Logger())

	agent.MustRegisterAction("ping", pingAction)

	return agents.Agent(agent), nil
}

func pingAction(ctx context.Context, req *mcorpc.Request, reply *mcorpc.Reply, agent *mcorpc.Agent, conn choria.ConnectorInfo) {
	i := &EchoRequest{}
	if !mcorpc.ParseRequestData(i, req, reply) {
		return
	}

	reply.Data = &EchoReply{Message: i.Message}
}

// ChoriaPlugin produces the Choria pluggable plugin it uses the metadata
// to dynamically answer questions of name and version
func ChoriaPlugin() plugin.Pluggable {
	return mcorpc.NewChoriaAgentPlugin(metadata, New)
}