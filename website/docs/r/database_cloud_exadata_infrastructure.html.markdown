---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_cloud_exadata_infrastructure"
sidebar_current: "docs-oci-resource-database-cloud_exadata_infrastructure"
description: |-
  Provides the Cloud Exadata Infrastructure resource in Oracle Cloud Infrastructure Database service
---

# oci_database_cloud_exadata_infrastructure
This resource provides the Cloud Exadata Infrastructure resource in Oracle Cloud Infrastructure Database service.

Creates a cloud Exadata infrastructure resource.

## Example Usage

```hcl
resource "oci_database_cloud_exadata_infrastructure" "test_cloud_exadata_infrastructure" {
	#Required
	availability_domain = var.cloud_exadata_infrastructure_availability_domain
	compartment_id = var.compartment_id
	display_name = var.cloud_exadata_infrastructure_display_name
	shape = var.cloud_exadata_infrastructure_shape

	#Optional
	compute_count = var.cloud_exadata_infrastructure_compute_count
	defined_tags = var.cloud_exadata_infrastructure_defined_tags
	freeform_tags = {"Department"= "Finance"}
	maintenance_window {
		#Required
		preference = var.cloud_exadata_infrastructure_maintenance_window_preference

		#Optional
		days_of_week {
			#Required
			name = var.cloud_exadata_infrastructure_maintenance_window_days_of_week_name
		}
		hours_of_day = var.cloud_exadata_infrastructure_maintenance_window_hours_of_day
		lead_time_in_weeks = var.cloud_exadata_infrastructure_maintenance_window_lead_time_in_weeks
		months {
			#Required
			name = var.cloud_exadata_infrastructure_maintenance_window_months_name
		}
		weeks_of_month = var.cloud_exadata_infrastructure_maintenance_window_weeks_of_month
	}
	storage_count = var.cloud_exadata_infrastructure_storage_count
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The availability domain where the cloud Exadata infrastructure is located.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compute_count` - (Optional) (Updatable) The number of compute servers for the cloud Exadata infrastructure.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - (Required) (Updatable) The user-friendly name for the cloud Exadata infrastructure resource. The name does not need to be unique. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `maintenance_window` - (Optional) (Updatable) 
	* `days_of_week` - (Optional) (Updatable) Days during the week when maintenance should be performed.
		* `name` - (Required) (Updatable) Name of the day of the week.
	* `hours_of_day` - (Optional) (Updatable) The window of hours during the day when maintenance should be performed. The window is a 4 hour slot. Valid values are
		* 0 - represents time slot 0:00 - 3:59 UTC - 4 - represents time slot 4:00 - 7:59 UTC - 8 - represents time slot 8:00 - 11:59 UTC - 12 - represents time slot 12:00 - 15:59 UTC - 16 - represents time slot 16:00 - 19:59 UTC - 20 - represents time slot 20:00 - 23:59 UTC
	* `lead_time_in_weeks` - (Optional) (Updatable) Lead time window allows user to set a lead time to prepare for a down time. The lead time is in weeks and valid value is between 1 to 4. 
	* `months` - (Optional) (Updatable) Months during the year when maintenance should be performed.
		* `name` - (Required) (Updatable) Name of the month of the year.
	* `preference` - (Required) (Updatable) The maintenance window scheduling preference.
	* `weeks_of_month` - (Optional) (Updatable) Weeks during the month when maintenance should be performed. Weeks start on the 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates, not days of the week. For example, to allow maintenance during the 2nd week of the month (from the 8th day to the 14th day of the month), use the value 2. Maintenance cannot be scheduled for the fifth week of months that contain more than 28 days. Note that this parameter works in conjunction with the  daysOfWeek and hoursOfDay parameters to allow you to specify specific days of the week and hours that maintenance will be performed. 
* `shape` - (Required) The shape of the cloud Exadata infrastructure resource. 
* `storage_count` - (Optional) (Updatable) The number of storage servers for the cloud Exadata infrastructure.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The name of the availability domain that the cloud Exadata infrastructure resource is located in.
* `available_storage_size_in_gbs` - The available storage can be allocated to the cloud Exadata infrastructure resource, in gigabytes (GB).
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compute_count` - The number of compute servers for the cloud Exadata infrastructure.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the cloud Exadata infrastructure resource. The name does not need to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Exadata infrastructure resource.
* `last_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last maintenance run.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `maintenance_window` - 
	* `days_of_week` - Days during the week when maintenance should be performed.
		* `name` - Name of the day of the week.
	* `hours_of_day` - The window of hours during the day when maintenance should be performed. The window is a 4 hour slot. Valid values are
		* 0 - represents time slot 0:00 - 3:59 UTC - 4 - represents time slot 4:00 - 7:59 UTC - 8 - represents time slot 8:00 - 11:59 UTC - 12 - represents time slot 12:00 - 15:59 UTC - 16 - represents time slot 16:00 - 19:59 UTC - 20 - represents time slot 20:00 - 23:59 UTC
	* `lead_time_in_weeks` - Lead time window allows user to set a lead time to prepare for a down time. The lead time is in weeks and valid value is between 1 to 4. 
	* `months` - Months during the year when maintenance should be performed.
		* `name` - Name of the month of the year.
	* `preference` - The maintenance window scheduling preference.
	* `weeks_of_month` - Weeks during the month when maintenance should be performed. Weeks start on the 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates, not days of the week. For example, to allow maintenance during the 2nd week of the month (from the 8th day to the 14th day of the month), use the value 2. Maintenance cannot be scheduled for the fifth week of months that contain more than 28 days. Note that this parameter works in conjunction with the  daysOfWeek and hoursOfDay parameters to allow you to specify specific days of the week and hours that maintenance will be performed. 
* `next_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the next maintenance run.
* `shape` - The model name of the cloud Exadata infrastructure resource. 
* `state` - The current lifecycle state of the cloud Exadata infrastructure resource.
* `storage_count` - The number of storage servers for the cloud Exadata infrastructure.
* `time_created` - The date and time the cloud Exadata infrastructure resource was created.
* `total_storage_size_in_gbs` - The total storage allocated to the cloud Exadata infrastructure resource, in gigabytes (GB).

## Import

CloudExadataInfrastructures can be imported using the `id`, e.g.

```
$ terraform import oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure "id"
```
