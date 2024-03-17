# Shopping-Complex-API
The Shopping Complex API provides endpoints to manage shops within a shopping complex. Users can create, read, update, and delete shops, as well as retrieve shops based on their rent payment status. Additionally, the API allows for updating a shop's status and name.

## API Documentation
[View API Documentation](https://documenter.getpostman.com/view/23129267/2sA2xnxA7D)

## Endpoints

- `/` (GET): Home endpoint
- `/getShops` (GET): Get all shops endpoint
- `/getShop/{id}` (GET): Get shop by ID endpoint
- `/createShop` (POST): Create a new shop endpoint
- `/updateStatus/{id}` (PUT): Update shop status by ID endpoint
- `/updateName/{id}` (PUT): Update shop name by ID endpoint
- `/deleteShop/{id}` (DELETE): Delete shop by ID endpoint
- `/statusTrue` (GET): Endpoint for shops with rent paid
- `/statusFalse` (GET): Endpoint for shops with rent not paid
