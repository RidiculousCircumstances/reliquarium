package interfaces

type ServiceName string

const (
	Eden        ServiceName = "eden"
	TraceFace   ServiceName = "trace_face"
	StreamMiner ServiceName = "stream_miner"
)

type CommandType string

const (
	Pause         CommandType = "pause"
	TakeSnapshots CommandType = "takeSnapshots"
	Resume        CommandType = "resume"
)

type ServiceRegistry interface {
	ConfirmServiceCommand(serviceName ServiceName, command CommandType) bool
	IsCommandConfirmed(command CommandType) bool
	GetUnconfirmedServices(command CommandType) []ServiceName
	ResetServiceStates()
}
