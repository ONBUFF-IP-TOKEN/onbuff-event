package proc

var gProcMgr *ProcManager

type ProcManager struct {
	Cmd *Cmd
}

func NewProcManager() *ProcManager {
	gProcMgr = new(ProcManager)

	gProcMgr.Cmd = NewCmd(gProcMgr)
	return gProcMgr
}

func GetProcManager() *ProcManager {
	return gProcMgr
}

func (o *ProcManager) Init() error {
	o.Cmd.StartCommand()
	return nil
}
