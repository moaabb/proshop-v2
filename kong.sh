#!/bin/bash

# Define the Kong Admin API URL
KONG_ADMIN_API="http://localhost:8001"

# Function to create a service
create_service() {
  local service_name="$1"
  local service_url="$2"
  
  echo "Creating service: $service_name"
  
  curl -i -X POST --url "$KONG_ADMIN_API/services/" \
    --data "name=$service_name" \
    --data "url=$service_url"
}

# Function to create a route for a service
create_route() {
  local route_name="$1"
  local service_name="$2"
  local route_path="$3"
  
  echo "Creating route: $route_name for service: $service_name"
  
  curl -i -X POST --url "$KONG_ADMIN_API/services/$service_name/routes" \
    --data "paths[]=$route_path"
}

# Define the services and routes
create_service "order-service" "http://172.21.193.94:8080"
create_route "order-route" "order-service" "/orders"

create_service "product-service" "http://172.21.193.94:5000"
create_route "product-route" "product-service" "/products"

create_service "user-service" "http://172.21.193.94:3000"
create_route "user-route" "user-service" "/users"

echo "Configuration completed."


curl -X POST http://localhost:8001/services/order-service/plugins \
   --data "name=cors"  \
   --data "config.origins=http://http://172.21.193.94:3000"  \
   --data "config.methods=GET"  \
   --data "config.methods=POST"  \
   --data "config.headers=Accept"  \
   --data "config.headers=Accept-Version"  \
   --data "config.headers=Content-Length"  \
   --data "config.headers=Content-MD5"  \
   --data "config.headers=Content-Type"  \
   --data "config.headers=Date"  \
   --data "config.headers=X-Auth-Token"  \
   --data "config.exposed_headers=X-Auth-Token"  \
   --data "config.credentials=true"  \
   --data "config.max_age=3600"

curl -X POST http://localhost:8001/services/product-service/plugins \
   --data "name=cors"  \
   --data "config.origins=http://http://172.21.193.94:3000"  \
   --data "config.methods=GET"  \
   --data "config.methods=POST"  \
   --data "config.headers=Accept"  \
   --data "config.headers=Accept-Version"  \
   --data "config.headers=Content-Length"  \
   --data "config.headers=Content-MD5"  \
   --data "config.headers=Content-Type"  \
   --data "config.headers=Date"  \
   --data "config.headers=X-Auth-Token"  \
   --data "config.exposed_headers=X-Auth-Token"  \
   --data "config.credentials=true"  \
   --data "config.max_age=3600"

curl -X POST http://localhost:8001/services/user-service/plugins \
   --data "name=cors"  \
   --data "config.origins=http://http://172.21.193.94:3000"  \
   --data "config.methods=GET"  \
   --data "config.methods=POST"  \
   --data "config.headers=Accept"  \
   --data "config.headers=Accept-Version"  \
   --data "config.headers=Content-Length"  \
   --data "config.headers=Content-MD5"  \
   --data "config.headers=Content-Type"  \
   --data "config.headers=Date"  \
   --data "config.headers=X-Auth-Token"  \
   --data "config.exposed_headers=X-Auth-Token"  \
   --data "config.credentials=true"  \
   --data "config.max_age=3600"
