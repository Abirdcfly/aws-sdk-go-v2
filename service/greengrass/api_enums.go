// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

// The current status of the bulk deployment.
type BulkDeploymentStatus string

// Enum values for BulkDeploymentStatus
const (
	BulkDeploymentStatusInitializing BulkDeploymentStatus = "Initializing"
	BulkDeploymentStatusRunning      BulkDeploymentStatus = "Running"
	BulkDeploymentStatusCompleted    BulkDeploymentStatus = "Completed"
	BulkDeploymentStatusStopping     BulkDeploymentStatus = "Stopping"
	BulkDeploymentStatusStopped      BulkDeploymentStatus = "Stopped"
	BulkDeploymentStatusFailed       BulkDeploymentStatus = "Failed"
)

func (enum BulkDeploymentStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum BulkDeploymentStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The type of deployment. When used for ''CreateDeployment'', only ''NewDeployment''
// and ''Redeployment'' are valid.
type DeploymentType string

// Enum values for DeploymentType
const (
	DeploymentTypeNewDeployment        DeploymentType = "NewDeployment"
	DeploymentTypeRedeployment         DeploymentType = "Redeployment"
	DeploymentTypeResetDeployment      DeploymentType = "ResetDeployment"
	DeploymentTypeForceResetDeployment DeploymentType = "ForceResetDeployment"
)

func (enum DeploymentType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DeploymentType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type EncodingType string

// Enum values for EncodingType
const (
	EncodingTypeBinary EncodingType = "binary"
	EncodingTypeJson   EncodingType = "json"
)

func (enum EncodingType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EncodingType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// Specifies whether the Lambda function runs in a Greengrass container (default)
// or without containerization. Unless your scenario requires that you run without
// containerization, we recommend that you run in a Greengrass container. Omit
// this value to run the Lambda function with the default containerization for
// the group.
type FunctionIsolationMode string

// Enum values for FunctionIsolationMode
const (
	FunctionIsolationModeGreengrassContainer FunctionIsolationMode = "GreengrassContainer"
	FunctionIsolationModeNoContainer         FunctionIsolationMode = "NoContainer"
)

func (enum FunctionIsolationMode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum FunctionIsolationMode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type LoggerComponent string

// Enum values for LoggerComponent
const (
	LoggerComponentGreengrassSystem LoggerComponent = "GreengrassSystem"
	LoggerComponentLambda           LoggerComponent = "Lambda"
)

func (enum LoggerComponent) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum LoggerComponent) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type LoggerLevel string

// Enum values for LoggerLevel
const (
	LoggerLevelDebug LoggerLevel = "DEBUG"
	LoggerLevelInfo  LoggerLevel = "INFO"
	LoggerLevelWarn  LoggerLevel = "WARN"
	LoggerLevelError LoggerLevel = "ERROR"
	LoggerLevelFatal LoggerLevel = "FATAL"
)

func (enum LoggerLevel) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum LoggerLevel) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type LoggerType string

// Enum values for LoggerType
const (
	LoggerTypeFileSystem    LoggerType = "FileSystem"
	LoggerTypeAwscloudWatch LoggerType = "AWSCloudWatch"
)

func (enum LoggerType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum LoggerType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The type of permission a function has to access a resource.
type Permission string

// Enum values for Permission
const (
	PermissionRo Permission = "ro"
	PermissionRw Permission = "rw"
)

func (enum Permission) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Permission) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The piece of software on the Greengrass core that will be updated.
type SoftwareToUpdate string

// Enum values for SoftwareToUpdate
const (
	SoftwareToUpdateCore     SoftwareToUpdate = "core"
	SoftwareToUpdateOtaAgent SoftwareToUpdate = "ota_agent"
)

func (enum SoftwareToUpdate) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SoftwareToUpdate) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The minimum level of log statements that should be logged by the OTA Agent
// during an update.
type UpdateAgentLogLevel string

// Enum values for UpdateAgentLogLevel
const (
	UpdateAgentLogLevelNone    UpdateAgentLogLevel = "NONE"
	UpdateAgentLogLevelTrace   UpdateAgentLogLevel = "TRACE"
	UpdateAgentLogLevelDebug   UpdateAgentLogLevel = "DEBUG"
	UpdateAgentLogLevelVerbose UpdateAgentLogLevel = "VERBOSE"
	UpdateAgentLogLevelInfo    UpdateAgentLogLevel = "INFO"
	UpdateAgentLogLevelWarn    UpdateAgentLogLevel = "WARN"
	UpdateAgentLogLevelError   UpdateAgentLogLevel = "ERROR"
	UpdateAgentLogLevelFatal   UpdateAgentLogLevel = "FATAL"
)

func (enum UpdateAgentLogLevel) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum UpdateAgentLogLevel) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The architecture of the cores which are the targets of an update.
type UpdateTargetsArchitecture string

// Enum values for UpdateTargetsArchitecture
const (
	UpdateTargetsArchitectureArmv6l  UpdateTargetsArchitecture = "armv6l"
	UpdateTargetsArchitectureArmv7l  UpdateTargetsArchitecture = "armv7l"
	UpdateTargetsArchitectureX8664   UpdateTargetsArchitecture = "x86_64"
	UpdateTargetsArchitectureAarch64 UpdateTargetsArchitecture = "aarch64"
)

func (enum UpdateTargetsArchitecture) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum UpdateTargetsArchitecture) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The operating system of the cores which are the targets of an update.
type UpdateTargetsOperatingSystem string

// Enum values for UpdateTargetsOperatingSystem
const (
	UpdateTargetsOperatingSystemUbuntu      UpdateTargetsOperatingSystem = "ubuntu"
	UpdateTargetsOperatingSystemRaspbian    UpdateTargetsOperatingSystem = "raspbian"
	UpdateTargetsOperatingSystemAmazonLinux UpdateTargetsOperatingSystem = "amazon_linux"
	UpdateTargetsOperatingSystemOpenwrt     UpdateTargetsOperatingSystem = "openwrt"
)

func (enum UpdateTargetsOperatingSystem) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum UpdateTargetsOperatingSystem) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
