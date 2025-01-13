# Cloud Resource Inventory Management System

This project simulates a Cloud Resource Inventory Management System, designed to manage and store cloud resource data for multiple customers. The system is built to be scalable and can accommodate a growing number of resources and customers.

## Objective

The goal of this project is to design and implement a backend service that enables users to manage cloud resources for multiple customers. The system will allow:
- Storing and retrieving cloud resource data for different customers.
- Efficient management of customer-specific cloud resource inventory.
- Scalable architecture to accommodate future growth.

## Key Features

- **Customer Management**: Each customer has a unique profile and their associated cloud resources.
- **Resource Management**: Users can create, update, and delete cloud resources (e.g., virtual machines, storage, databases).
- **Scalable Design**: The system is designed to scale horizontally to handle a large number of customers and resources.
- **API Endpoints**: RESTful APIs for managing resources, including CRUD operations.

## Technologies Used

- **Go (Golang)**: Backend development with Golang.
- **gRPC**: For efficient communication and resource retrieval.
- **Docker & Docker Compose**: Containerization for easier setup and deployment.
- **Redis**: Used for event-driven communication and pub/sub.
- **PostgreSQL**: Relational database to store customer and resource data.
- **Asyncmon**: For event handling and message queue management.

## Requirements

- Docker
- Docker Compose
- Go 1.23 or above
- PostgreSQL

## Setup and Installation

### Step 1: Clone the Repository

Clone the repository to your local machine:

```bash
git clone https://github.com/azeezlala/assessment.git
cd assessment
```

### Step 2: Setup Environment Variables

Create an `.env` file in the root directory and add necessary environment variables for your database and services:

```env
PORT=8082
DB_PORT=5432
DB_USER=username
DB_PASSWORD=password
DB_NAME=cloud_inventory
GRPC_PORT=20012
REDIS_URL=localhost:6379
```

### Step 3: Start the server
In root directory, open your terminal and run 
```
  docker-compose up
```

This command will:
1. Build the services.
2. Start the PostgreSQL database, the API service, and other necessary containers.

The system will be running on the following ports:
- **API Service**: `localhost:8082` (for RESTful communication)
- **gRPC Service**: `localhost:20012` (for retrieving resources via gRPC)

### Step 4: Test the API
The API provides CRUD operations for managing customers and their cloud resources.

1. **Add a Customer**    
To create a new customer, send a POST request to the /customers endpoint:
```bash
  curl -X POST http://localhost:8082/customers -d '{"name": "Customer A", "email": "customerA@example.com"}' -H "Content-Type: application/json"
```

2. **Get Customer Details**.  
To retrieve customer details, send a GET request to /customers/{id}
```bash
  curl http://localhost:8082/customers/{id}
```
Replace {id} with the actual customer ID.

3. **Fetch Available Resorce**.  
this resources were already seeder
```bash
  http://localhost:8082/resources
```

4. **Assign a Cloud Resource to a Customer**.
```bash
  curl -X POST http://localhost:8082/resources -d '{"customer_id": "7bd7b6af-0128-4914-8e85-829bda009d1f", "resource_id":"328e1de8-a348-46e5-a0c3-884912344c14"}' -H "Content-Type: application/json"
```
the customer id represent the customer that was initially created, same for the resource_id

### Step 5: Use gRPC to Notification
You can also interact with the Notification Service using gRPC, you can fetch notifications from the notification service
To test gRPC, you can use the following command:
```bash 
  grpcurl -plaintext localhost:50051 list
```
This will list all available gRPC services and methods.

### Step 6: Stopping the Services
To stop the services, use the following command:
```bash
  docker-compose down
```
This will stop and remove the containers for the services.

### Events
The **Notification Service** is responsible for handling notifications related to the system's cloud resources. It subscribes to events and makes these notifications available to clients via gRPC.

#### Events Published:
1. **Customer Added Event**: When a new customer is added to the system, an event is published, and a notification is triggered.
2. **Resource Added Event**: When a new resource is added to the system, an event is published, and a notification is triggered.

These events are published to the **Notification Service**, which subscribes to them and makes the notifications available via gRPC. Clients can fetch the notifications using the gRPC service.

### Testing the Notification Service
You can test the system using the following API endpoints:  
- **Fetch Notification by ID** (GET):  
`http://localhost:8082/notifications/{id}`  
This endpoint allows you to fetch a specific notification using its ID, you can pass the customer id as the id to fetch notification relating to customer creation

### Database Migration

The system uses automated database migrations to keep the schema up-to-date across different environments. This ensures that the database structure remains consistent and properly managed throughout the development lifecycle.

#### Steps to Run Migrations

1. **Migration Command**:
   You can run migrations with the following command:
   ```bash
   go run cmd/migrate/main.go migrate 
   ```
2. **Seeder Command**:
   You can run the seeder using this command:
   ```bash
   go run cmd/migrate/main.go seed 
   ```

#### Steps to Run Seeder

### CI Workflow

The project includes an automated Continuous Integration (CI) workflow to ensure consistent quality and streamline the development process. The CI workflow is managed using GitHub Actions, and it automatically runs tests, builds the Docker containers, and ensures everything is functioning as expected with every push.

#### Key CI Actions:
1. **Linting & Testing**: Each commit triggers linting and testing to ensure code quality and catch any potential errors early in the development process.
2. **Build Docker Containers**: The Docker containers for both the **API Service** and **Notification Service** are built and validated to ensure they are properly containerized.
3. **Deployment Preparation**: The workflow prepares for deployment, ensuring all necessary artifacts are ready to be pushed to the production environment.

This CI process runs on every push and pull request, so developers can be confident that the system is always up-to-date and error-free.

### Workflow Trigger:
- The workflow is triggered by every **push** and **pull request** to the repository.

#### Example CI Workflow in `.github/workflows/ci.yml`:

### Architecture
The system is composed of several microservices:

- **API Service:** A RESTful service that allows users to manage customer and resource data.
- **Asyncmon:** For asynchronous event handling, such as notifying other services when a resource is added or updated.
- **GRPC:** Used in the notification service to fetch notifications

### Conclusion
This project demonstrates a scalable backend system for managing cloud resources for multiple customers. It uses RESTful APIs for basic interaction, gRPC for faster data retrieval, and PostgreSQL for data storage. You can test the system by interacting with the API and using gRPC for fetching notifications.

