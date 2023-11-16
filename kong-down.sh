#!/bin/bash

# Define the Kong Admin API URL
KONG_ADMIN_API="http://localhost:8001"

# Function to delete a service
delete_service() {
  local service_name="$1"
  
  echo "Deleting service: $service_name"
  
  curl -i -X DELETE --url "$KONG_ADMIN_API/services/$service_name"
}

# Function to delete a route for a service
delete_route() {
  local route_name="$1"
  
  echo "Deleting route: $route_name"/services/{service name or id}/routes/{route name or id}

  
  curl -i -X DELETE --url "$KONG_ADMIN_API/routes/$route_name"
}

# Delete services and routes
delete_route "order-route"
delete_service "order-service"

delete_route "product-route"
delete_service "product-service"

delete_route "user-route"
delete_service "user-service"

delete_route "auth-route"
delete_service "auth-service"

# Delete CORS plugin
echo "Deleting CORS plugin configuration"
curl -i -X DELETE --url "$KONG_ADMIN_API/plugins/cors"

echo "Deletion completed."
    