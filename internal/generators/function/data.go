package function

type Data struct {
	Name                string
	Source              string
	Runtime             string
	SourceArchiveBucket string
	SourceArchiveObject string
	TriggerHTTP         string
	EntryPoint          string
	Envars              map[string]string
}
