package resources

import "strings"

type ResourceGCS struct {
	Type  string
	Name  string
	Label string
}

func ParseResourceGCS(gcs string, labels []string) *ResourceGCS {
	var gcsType, name, label string

	switch {
	case strings.HasPrefix(gcs, "gs://"):
		gcsType, name, label = parseGCSByGS(gcs)
	case strings.HasPrefix(gcs, "http"):
		gcsType, name, label = parseGCSByURL(gcs)
	default:
		gcsType, name, label = parseGCSGeneric(gcs)
	}

	if gcsType == "" && len(labels) > 0 {
		gcsType = labels[0]
	}

	if label == "" && len(labels) > 1 {
		label = labels[1]
	}

	return &ResourceGCS{Type: gcsType, Name: name, Label: label}
}

func parseGCSByGS(gcs string) (gcsType, name, label string) {
	parts := strings.Split(gcs, "gs://")

	if len(parts) > 1 {
		parts = strings.Split(parts[1], "/")

		if len(parts) > 0 {
			parts = strings.Split(parts[0], ".")

			if strings.HasPrefix(parts[0], "google_") {
				gcsType = parts[0]
				label = parts[1]
			} else {
				name = parts[0]
			}
		}
	}

	return
}

func parseGCSByURL(gcs string) (gcsType, name, label string) {
	parts := strings.Split(gcs, "//")

	if len(parts) > 1 {
		parts = strings.Split(parts[1], "/")

		if len(parts) > 0 {
			name = parts[len(parts)-1]
		}
	}

	return
}

func parseGCSGeneric(gcs string) (gcsType, name, label string) {
	parts := strings.Split(gcs, ".")

	if len(parts) > 1 && strings.HasPrefix(parts[0], "google_") {
		gcsType = parts[0]
		label = parts[1]
	} else {
		name = gcs
	}

	return gcsType, name, label
}
