package drawiotoresources

import (
	"regexp"

	drawioxml "github.com/joselitofilho/drawio-parser-go/pkg/parser/xml"

	"github.com/joselitofilho/gcp-terraform-generator/internal/resources"
)

// Transform parses resources from the MxFile.
func Transform(mxFile *drawioxml.MxFile) (*resources.ResourceCollection, error) {
	resc := resources.NewResourceCollection()

	for i := range mxFile.Diagram.MxGraphModel.Root.MxCells {
		cell := mxFile.Diagram.MxGraphModel.Root.MxCells[i]

		resource := createResource(cell.ID, cell.Value, cell.Style)
		if resource != nil {
			resc.AddResource(resource)
		}
	}

	resourcesMap := map[string]resources.Resource{}
	for _, resource := range resc.Resources {
		resourcesMap[resource.ID()] = resource
	}

	for i := range mxFile.Diagram.MxGraphModel.Root.MxCells {
		cell := mxFile.Diagram.MxGraphModel.Root.MxCells[i]
		if cell.Source != "" && cell.Target != "" {
			source := resourcesMap[cell.Source]
			target := resourcesMap[cell.Target]

			if source != nil && target != nil {
				resc.AddRelationship(source, target)
			}
		}
	}

	return resc, nil
}

// createResource creates a resource based on cell data.
func createResource(id, value, style string) resources.Resource {
	// NOTE How to identify resource:
	// - PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHhtbG5zOnY9Imh0dHBzOi8vdmVjdGEuaW8vbmFubyIgd2lkdGg9Ij
	// - followed by 72 chars
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
		return resources.NewGenericResource(id, value, resources.AppEngine)
	case reBigQuery.MatchString(style):
		return resources.NewGenericResource(id, value, resources.BigQuery)
	case reCloudBigTable.MatchString(style):
		return resources.NewGenericResource(id, value, resources.CloudBigTable)
	case reCloudFunction.MatchString(style):
		return resources.NewGenericResource(id, value, resources.CloudFunction)
	case reCloudStorage.MatchString(style):
		return resources.NewGenericResource(id, value, resources.CloudStorage)
	case reDataflow.MatchString(style):
		return resources.NewGenericResource(id, value, resources.Dataflow)
	case resIoTCore.MatchString(style):
		return resources.NewGenericResource(id, value, resources.IoTCore)
	case resPubSub.MatchString(style):
		return resources.NewGenericResource(id, value, resources.PubSub)
	default:
		return nil
	}
}
