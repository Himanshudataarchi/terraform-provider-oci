package monitoring

import (
	oci_monitoring "github.com/oracle/oci-go-sdk/v65/monitoring"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("monitoring", monitoringResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportMonitoringAlarmHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_monitoring_alarm",
	DatasourceClass:        "oci_monitoring_alarms",
	DatasourceItemsAttr:    "alarms",
	ResourceAbbreviation:   "alarm",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_monitoring.AlarmLifecycleStateActive),
	},
}

var monitoringResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportMonitoringAlarmHints},
	},
}
