package resourcestoyaml

import "github.com/joselitofilho/gcp-terraform-generator/internal/resources"

func (t *Transformer) buildIoTCoreToPubSub(core, pubSub resources.Resource) {
	coreID := core.ID()
	t.pubSubByIoTCoreID[coreID] = append(t.pubSubByIoTCoreID[coreID], pubSub)
}

func (t *Transformer) buildPubSubToDataFlow(pubSub, dataFlow resources.Resource) {
	dataFlowID := dataFlow.ID()
	t.pubSubByDataFlowID[dataFlowID] = append(t.pubSubByDataFlowID[dataFlowID], pubSub)
}
