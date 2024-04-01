package resourcestoyaml

import "github.com/joselitofilho/gcp-terraform-generator/internal/resources"

func (t *Transformer) buildIoTCoreToPubSub(core, pubsub resources.Resource) {
	coreID := core.ID()
	t.pubSubByIoTCoreID[coreID] = append(t.pubSubByIoTCoreID[coreID], pubsub)
}
