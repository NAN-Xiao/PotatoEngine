package agent

import "container/list"

type AgentMgr struct {
	_init   bool
	_agents *list.List
}

//取出底部元素 先入先出
func (mgr *AgentMgr) Pop() *Agent {
	value := mgr._agents.Front()
	if value == nil {
		return nil
	}
	agent := value.Value.(*Agent)
	if agent == nil {
		return nil
	}
	return agent
}

//添加agent
func (mgr *AgentMgr) AddAgent(agent *Agent) {
	mgr._agents.PushBack(agent)
}

var instance *AgentMgr

func GetAgentMgr() *AgentMgr {
	if instance == nil {
		instance = &AgentMgr{
			_init:   true,
			_agents: list.New(),
		}
	}
	return instance
}
