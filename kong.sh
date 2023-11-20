#!/bin/bash

# Define the Kong Admin API URL
DOMAIN="10.0.0.9"
KONG_ADMIN_API="http://$DOMAIN:8001"

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
    --data "paths[]=$route_path" \
    --data "name=$route_name"
}

# Define the services and routes
create_service "order-service" "http://$DOMAIN:8080"
create_route "order-route" "order-service" "/orders"

create_service "product-service" "http://$DOMAIN:5000"
create_route "product-route" "product-service" "/products"

create_service "user-service" "http://$DOMAIN:5050"
create_route "user-route" "user-service" "/users"

create_service "auth-service" "http://$DOMAIN:2020"
create_route "auth-route" "auth-service" "/auth"

create_service "asset-service" "http://$DOMAIN:6060"
create_route "asset-route" "asset-service" "/uploads"

echo "Creating CORS plugin"

curl -X POST http://$DOMAIN:8001/plugins \
   --data "name=cors"  \
   --data "config.origins=http://$DOMAIN:3000"  \
   --data "config.origins=http://localhost:3000"  \
   --data "config.methods=GET"  \
   --data "config.methods=POST"  \
   --data "config.methods=PATCH"  \
   --data "config.methods=PUT"  \
   --data "config.methods=DELETE"  \
   --data "config.methods=OPTIONS"  \
   --data "config.headers=Accept"  \
   --data "config.headers=Accept-Version"  \
   --data "config.headers=Content-Length"  \
   --data "config.headers=Content-MD5"  \
   --data "config.headers=Content-Type"  \
   --data "config.headers=Authorization"  \
   --data "config.headers=Date"  \
   --data "config.headers=X-Auth-Token"  \
   --data "config.exposed_headers=X-Auth-Token"  \
   --data "config.exposed_headers=Access-Control-Allow-Origin"  \
   --data "config.credentials=true"  \
   --data "config.max_age=3600"

echo "Configuration completed."

