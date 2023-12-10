import typing
from grafanalib.core import (
    Alert, AlertCondition, Dashboard, Graph,
    GreaterThan, OP_AND, OPS_FORMAT, Row, RTYPE_SUM, SECONDS_FORMAT,
    SHORT_FORMAT, single_y_axis, Target, TimeRange, YAxes, YAxis, DataSourceInput, MILLISECONDS_FORMAT
)

# ------------------------------
# For Gostats we recommend using
# Grafana dashboard ID 10826 - https://grafana.com/grafana/dashboards/10826
#

DATASOURCE_NAME = "DS_PROM"
DATASOURCE = "${%s}" % DATASOURCE_NAME


class NebulaAdmin(object):
    APIS = [
        "create_execution",
        "create_launch_plan",
        "create_task",
        "create_workflow",
        "create_node_execution_event",
        "create_task_execution_event",
        "get_execution",
        "get_launch_plan",
        "get_task",
        "get_workflow",
        "get_node_execution",
        "get_task_execution",
        "get_active_launch_plan",
        "list_execution",
        "list_launch_plan",
        "list_task",
        "list_workflow",
        "list_node_execution",
        "list_task_execution",
        "list_active_launch_plans",
    ]

    ENTITIES = [
        "executions",
        "task_executions",
        "node_executions",
        "workflows",
        "launch_plans",
        "project",
    ]

    DB_OPS = [
        "get",
        "list",
        "create",
        "update",
        "list_identifiers",
        "delete",
        "exists",
    ]

    @staticmethod
    def error_codes(api: str, interval: int = 1) -> Graph:
        return Graph(
            title=f"{api} return codes",
            dataSource=DATASOURCE,
            targets=[
                Target(
                    expr=f'sum(irate(nebula:admin:admin:{api}:codes:OK[{interval}m]))',
                    legendFormat="ok",
                    refId='A',
                ),
                Target(
                    expr=f'sum(irate(nebula:admin:admin:{api}:codes:InvalidArgument[{interval}m]))',
                    legendFormat="invalid-args",
                    refId='B',
                ),
                Target(
                    expr=f'sum(irate(nebula:admin:admin:{api}:codes:AlreadyExists[{interval}m]))',
                    legendFormat="already-exists",
                    refId='C',
                ),
                Target(
                    expr=f'sum(irate(nebula:admin:admin:{api}:codes:FailedPrecondition[{interval}m]))',
                    legendFormat="failed-precondition",
                    refId='D',
                ),
            ],
            yAxes=YAxes(
                YAxis(format=OPS_FORMAT),
                YAxis(format=SHORT_FORMAT),
            ),
        )

    @staticmethod
    def error_vs_success(api: str, interval: int = 1) -> Graph:
        return Graph(
            title=f"{api} success vs errors",
            dataSource=DATASOURCE,
            targets=[
                Target(
                    expr=f'sum(irate(nebula:admin:admin:{api}:errors[{interval}m]))',
                    legendFormat="errors",
                    refId='A',
                ),
                Target(
                    expr=f'sum(irate(nebula:admin:admin:{api}:success[{interval}m]))',
                    legendFormat="success",
                    refId='B',
                ),
            ],
            yAxes=YAxes(
                YAxis(format=OPS_FORMAT),
                YAxis(format=SHORT_FORMAT),
            ),
        )

    @staticmethod
    def api_latency(api: str, interval: int = 1) -> Graph:
        return Graph(
            title=f"{api} Latency",
            dataSource=DATASOURCE,
            targets=[
                Target(
                    expr=f'max(nebula:admin:admin:{api}:duration_ms) by (quantile)',
                    refId='A',
                ),
            ],
            yAxes=single_y_axis(format=MILLISECONDS_FORMAT),
        )

    @staticmethod
    def create_api_row(api: str, collapse: bool, interval: int = 1) -> Row:
        return Row(
            title=f"{api} stats",
            collapse=collapse,
            panels=[
                NebulaAdmin.error_codes(api, interval),
                NebulaAdmin.error_vs_success(api, interval),
                NebulaAdmin.api_latency(api, interval),
            ]
        )

    @staticmethod
    def db_latency(entity: str, op: str, interval: int = 1) -> Graph:
        return Graph(
            title=f"{op} Latency",
            dataSource=DATASOURCE,
            targets=[
                Target(
                    expr=f'sum(nebula:admin:admin:database:{entity}:{op}_ms) by (quantile)',
                    refId='A',
                ),
            ],
            yAxes=single_y_axis(format=MILLISECONDS_FORMAT),
        )

    @staticmethod
    def create_entity_db_row_latency(entity: str, collapse: bool, interval: int = 1) -> Row:
        r = Row(
            title=f"DB {entity} ops stats",
            collapse=collapse,
            panels=[],
        )
        for op in NebulaAdmin.DB_OPS:
            r.panels.append(NebulaAdmin.db_latency(entity, op=op, interval=interval))
        return r

    @staticmethod
    def db_count(entity: str, op: str, interval: int = 1) -> Graph:
        return Graph(
            title=f"{op} Count Ops",
            dataSource=DATASOURCE,
            targets=[
                Target(
                    expr=f'sum(rate(nebula:admin:admin:database:{entity}:{op}_ms_count[{interval}m]))',
                    refId='A',
                ),
            ],
            yAxes=YAxes(
                YAxis(format=OPS_FORMAT),
                YAxis(format=SHORT_FORMAT),
            ),
        )

    @staticmethod
    def create_entity_db_count(entity: str, collapse: bool, interval: int = 1) -> Row:
        r = Row(
            title=f"DB {entity} ops stats",
            collapse=collapse,
            panels=[],
        )
        for op in NebulaAdmin.DB_OPS:
            r.panels.append(NebulaAdmin.db_count(entity, op=op, interval=interval))
        return r

    @staticmethod
    def create_all_entity_db_rows(collapse: bool, interval: int = 1) -> typing.List[Row]:
        rows = []
        for entity in NebulaAdmin.ENTITIES:
            rows.append(NebulaAdmin.create_entity_db_row_latency(entity=entity, collapse=collapse, interval=interval))
            rows.append(NebulaAdmin.create_entity_db_count(entity=entity, collapse=collapse, interval=interval))
        return rows

    @staticmethod
    def create_all_apis(interval: int = 5) -> typing.List[Row]:
        rows = []
        for api in NebulaAdmin.APIS:
            rows.append(NebulaAdmin.create_api_row(api, collapse=True, interval=interval))
        return rows

    @staticmethod
    def create_all_rows(interval: int = 5) -> typing.List[Row]:
        rows = []
        rows.extend(NebulaAdmin.create_all_entity_db_rows(collapse=True, interval=interval))
        rows.extend(NebulaAdmin.create_all_apis(interval))
        return rows


dashboard = Dashboard(
    tags=["nebula", "prometheus", "nebulaadmin", "nebula-controlplane"],
    inputs=[
        DataSourceInput(
            name=DATASOURCE_NAME,
            label="Prometheus",
            description="Prometheus server that connects to Nebula",
            pluginId="prometheus",
            pluginName="Prometheus",
        ),
    ],
    editable=False,
    title="NebulaAdmin Dashboard (via Prometheus)",
    rows=NebulaAdmin.create_all_rows(),
    description="NebulaAdmin/Control Plane Dashboard. This is great for monitoring NebulaAdmin and the Service API.",
).auto_panel_ids()

if __name__ == "__main__":
    print(dashboard.to_json_data())
