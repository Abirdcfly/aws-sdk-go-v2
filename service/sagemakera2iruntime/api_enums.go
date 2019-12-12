// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemakera2iruntime

type ContentClassifier string

// Enum values for ContentClassifier
const (
	ContentClassifierFreeOfPersonallyIdentifiableInformation ContentClassifier = "FreeOfPersonallyIdentifiableInformation"
	ContentClassifierFreeOfAdultContent                      ContentClassifier = "FreeOfAdultContent"
)

func (enum ContentClassifier) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ContentClassifier) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type HumanLoopStatus string

// Enum values for HumanLoopStatus
const (
	HumanLoopStatusInProgress HumanLoopStatus = "InProgress"
	HumanLoopStatusFailed     HumanLoopStatus = "Failed"
	HumanLoopStatusCompleted  HumanLoopStatus = "Completed"
	HumanLoopStatusStopped    HumanLoopStatus = "Stopped"
	HumanLoopStatusStopping   HumanLoopStatus = "Stopping"
)

func (enum HumanLoopStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum HumanLoopStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAscending  SortOrder = "Ascending"
	SortOrderDescending SortOrder = "Descending"
)

func (enum SortOrder) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SortOrder) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
