# JSON SCHEMA

This folder contains the JSON Schema of execution event, and this schema will
be used by [cloudevent publisher](https://github.com/nebulaclouds/nebulaadmin/blob/a0ca4b07d3b1c3cccfe3830307df50bc73152ddb/pkg/async/cloudevent/implementations/cloudevent_publisher.go#L96).

The URL for the JSON file will be included with the cloudevent message. For example
```json
{
    "specversion" : "1.0",
    "type" : "com.nebula.resource.workflow",
    "source" : "https://github.com/nebulaclouds/nebulaadmin",
    "id" : "D234-1234-1234",
    "time" : "2018-04-05T17:31:00Z",
    "jsonschemaurl": "https://github.com/nebulaclouds/nebulaidl/blob/master/jsonschema/workflow_execution.json",
    "data" : "workflow execution event"
}
```