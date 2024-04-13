package resourcestoyaml

import "github.com/diagram-code-generator/resources/pkg/resources"

func (t *Transformer) buildAppEngineToBigTable(appEngine, bigTable resources.Resource) {
	bigTableID := bigTable.ID()
	t.appEngineByBigTableID[bigTableID] = append(t.appEngineByBigTableID[bigTableID], appEngine)
}

func (t *Transformer) buildDataFlowToBigQuery(dataFlow, bq resources.Resource) {
	dataFlowID := dataFlow.ID()
	t.bqTablesByDataFlowID[dataFlowID] = append(t.bqTablesByDataFlowID[dataFlowID], bq)
}

func (t *Transformer) buildDataFlowToPubSub(dataFlow, pubSub resources.Resource) {
	dataFlowID := dataFlow.ID()
	t.outputPubSubByDataFlowID[dataFlowID] = append(t.outputPubSubByDataFlowID[dataFlowID], pubSub)
}

func (t *Transformer) buildDataFlowToStorage(dataFlow, storage resources.Resource) {
	dataFlowID := dataFlow.ID()
	t.storageByDataFlowID[dataFlowID] = append(t.storageByDataFlowID[dataFlowID], storage)
}

func (t *Transformer) buildFunctionToPubSub(function, pubSub resources.Resource) {
	pubSubID := pubSub.ID()
	t.functionPublisherByPubSubID[pubSubID] = append(t.functionPublisherByPubSubID[pubSubID], function)

	functionID := function.ID()
	t.pubSubsFromFunctionID[functionID] = append(t.pubSubsFromFunctionID[functionID], pubSub)
}

func (t *Transformer) buildIoTCoreToPubSub(core, pubSub resources.Resource) {
	coreID := core.ID()
	t.pubSubByIoTCoreID[coreID] = append(t.pubSubByIoTCoreID[coreID], pubSub)
}

func (t *Transformer) buildPubSubToAppEngine(pubSub, appEngine resources.Resource) {
	pubSubID := pubSub.ID()
	t.appEngineByPubSubID[pubSubID] = append(t.appEngineByPubSubID[pubSubID], appEngine)
}

func (t *Transformer) buildPubSubToFunction(pubSub, function resources.Resource) {
	pubSubID := pubSub.ID()
	t.functionSubscriberByPubSubID[pubSubID] = function

	functionID := function.ID()
	t.pubSubsToFunctionID[functionID] = append(t.pubSubsToFunctionID[functionID], pubSub)
}

func (t *Transformer) buildPubSubToDataFlow(pubSub, dataFlow resources.Resource) {
	dataFlowID := dataFlow.ID()
	t.inputPubSubByDataFlowID[dataFlowID] = append(t.inputPubSubByDataFlowID[dataFlowID], pubSub)
}
