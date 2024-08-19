package iintentqueues

type iSession struct {
	context *iContext
	home    string //the user home dir path

}

func (inst *iSession) openEndpoint() *iQueueEP {

	ep := new(iQueueEP)
	ep.session = inst
	ep.id = inst.context.nextQueueID()

	ep.open()

	inst.context.putEndpoint(ep)
	return ep
}
