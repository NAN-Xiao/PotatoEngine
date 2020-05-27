package agent

type AgentMgr struct {
	_init   bool
	_agents map[uint32]*Agent
}

//添加agent
func (this *AgentMgr) AddAgent(agent *Agent) {
	id := agent._cid
	_, ok := this._agents[id]
	if ok == true {
		return
	}
	this._agents[id] = agent
}

//得到注册的anent
func (this *AgentMgr) GetAgent(cid uint32) *Agent {
	v, ok := this._agents[cid]
	if ok {
		return v
	}
	return nil
}

func (this *AgentMgr) RemoveAgent(cid uint32) {
	_, ok := this._agents[cid]
	if ok {
		delete(this._agents, cid)
	}
}

var instance *AgentMgr

func GetAgentMgr() *AgentMgr {
	if instance == nil {
		instance = &AgentMgr{
			_init:   true,
			_agents: make(map[uint32]*Agent),
		}
	}
	return instance
}
