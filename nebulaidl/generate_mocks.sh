#!/bin/bash
set -e
set -x

mockery -dir=gen/pb-go/nebulaidl/service/ -all -output=clients/go/admin/mocks
mockery -dir=gen/pb-go/nebulaidl/datacatalog/ -name=DataCatalogClient -output=clients/go/datacatalog/mocks
