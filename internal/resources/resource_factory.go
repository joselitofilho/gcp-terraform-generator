package resources

import (
	"regexp"

	"github.com/diagram-code-generator/resources/pkg/resources"
)

type GCPResourceFactory struct{}

// CreateResource creates a resource based on cell data.
func (f *GCPResourceFactory) CreateResource(id, value, style string) resources.Resource {
	// NOTE How to identify resource:
	// - PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHhtbG5zOnY9Imh0dHBzOi8vdmVjdGEuaW8vbmFubyIgd2lkdGg9Ij
	// - followed by 72 chars.
	reAppEngine := regexp.MustCompile("IwIiBoZWlnaHQ9IjE2LjAyMDAwMDQ1Nzc2MzY3MiIgZmlsbC1ydWxlPSJldmVub2RkIiB2aW")
	reBigQuery := regexp.MustCompile("IwLjAwMTA0NTIyNzA1MDc4IiBoZWlnaHQ9IjIwLjAwMTA0NTIyNzA1MDc4IiBmaWxsLXJ1bG")
	reCloudBigTable := regexp.MustCompile("E3Ljk1Njk3Nzg0NDIzODI4IiBoZWlnaHQ9IjIwLjAwOTI1NjM2MjkxNTA0IiB2aWV3Qm94PS")
	reCloudFunction := regexp.MustCompile("IwIiBoZWlnaHQ9IjE5Ljk4OTk5OTc3MTExODE2NCIgdmlld0JveD0iMCAwIDIwIDE5Ljk4OT")
	reCloudStorage := regexp.MustCompile("IwIiBoZWlnaHQ9IjE2IiB2aWV3Qm94PSIwIDAgMjAgMTYiPiYjeGE7CTxzdHlsZSB0eXBlPS")
	reDataflow := regexp.MustCompile("E0LjUxOTk5OTUwNDA4OTM1NSIgaGVpZ2h0PSIyMCIgdmlld0JveD0iMCAwIDE0LjUxOTk5OT")
	resIoTCore := regexp.MustCompile("E5LjcwNjUzMTUyNDY1ODIwMyIgaGVpZ2h0PSIxOS45ODM4MjE4Njg4OTY0ODQiIGZpbGwtcn")
	resPubSub := regexp.MustCompile("E4LjMxOTk5OTY5NDgyNDIyIiBoZWlnaHQ9IjIwLjAwMDAwMTkwNzM0ODYzMyIgdmlld0JveD")

	switch {
	case reAppEngine.MatchString(style):
		return resources.NewGenericResource(id, value, AppEngine.String())
	case reBigQuery.MatchString(style):
		return resources.NewGenericResource(id, value, BigQuery.String())
	case reCloudBigTable.MatchString(style):
		return resources.NewGenericResource(id, value, BigTable.String())
	case reCloudFunction.MatchString(style):
		return resources.NewGenericResource(id, value, Function.String())
	case reCloudStorage.MatchString(style):
		return resources.NewGenericResource(id, value, Storage.String())
	case reDataflow.MatchString(style):
		return resources.NewGenericResource(id, value, Dataflow.String())
	case resIoTCore.MatchString(style):
		return resources.NewGenericResource(id, value, IoTCore.String())
	case resPubSub.MatchString(style):
		return resources.NewGenericResource(id, value, PubSub.String())
	default:
		return nil
	}
}
