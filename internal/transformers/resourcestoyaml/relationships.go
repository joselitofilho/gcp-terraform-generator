package resourcestoyaml

import "github.com/joselitofilho/gcp-terraform-generator/internal/resources"

func (t *Transformer) buildDataFlowToBigQuery(dataFlow, bq resources.Resource) {
	dataFlowID := dataFlow.ID()
	t.bqTablesByDataFlowID[dataFlowID] = append(t.bqTablesByDataFlowID[dataFlowID], bq)
}

func (t *Transformer) buildDataFlowToStorage(dataFlow, storage resources.Resource) {
	dataFlowID := dataFlow.ID()
	t.storageByDataFlowID[dataFlowID] = append(t.storageByDataFlowID[dataFlowID], storage)
}

func (t *Transformer) buildIoTCoreToPubSub(core, pubSub resources.Resource) {
	coreID := core.ID()
	t.pubSubByIoTCoreID[coreID] = append(t.pubSubByIoTCoreID[coreID], pubSub)
}

func (t *Transformer) buildPubSubToDataFlow(pubSub, dataFlow resources.Resource) {
	dataFlowID := dataFlow.ID()
	t.pubSubByDataFlowID[dataFlowID] = append(t.pubSubByDataFlowID[dataFlowID], pubSub)
}
