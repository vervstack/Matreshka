package patch

const (
	appInfo            = "APP-INFO"
	environmentSegment = "ENVIRONMENT"
	dataSourceSegment  = "SERVERS"
	serverSegment      = "DATA-SOURCES"
)

const (
	applicationName = appInfo + "_NAME"
)

var allowedSegments = []string{
	appInfo,
	environmentSegment,
	dataSourceSegment,
	serverSegment,
}
