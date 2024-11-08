## High-Level Overview

This Go application is designed as a high-throughput REST service capable of handling up to 10,000 requests per second. It features a single GET endpoint (/api/verve/accept) that processes an integer id as a mandatory parameter (for tracking unique requests) and an optional endpoint parameter (for sending unique request counts). The application logs the count of unique requests every minute and, if an endpoint is provided, sends this count to the specified HTTP endpoint.

In addition, the application includes extensions for switching to HTTP POST requests, ensuring unique ID tracking across load-balanced instances, and streaming unique counts to a distributed system.
Design and Implementation Approach
1. High-Throughput Design

To achieve high performance, the following strategies are implemented:

    Stateless Processing: Each request is processed independently, allowing the system to handle high volumes without complex state management.
    Concurrency: Go’s lightweight goroutines enable concurrent processing for tasks like logging and sending HTTP requests without blocking the main execution flow.
    In-Memory Data Store: A sync.Map is used for tracking unique IDs, providing thread-safe access and updates with low overhead.

2. Endpoint and Request Handling

   GET Endpoint: The /api/verve/accept endpoint accepts two parameters, id and endpoint, and responds with "ok" on successful processing.
   Error Handling: If id is missing, the server responds with a 400 Bad Request.
   Concurrency for Optional Endpoint: When the optional endpoint parameter is provided, a goroutine is launched to asynchronously send a unique count to the specified endpoint.

3. Periodic Unique Request Logging

   Scheduled Logging: Using time.Ticker, the application logs the count of unique requests every minute.
   Thread-Safe Reset: After logging, the sync.Map storing unique requests is reset for the new minute, ensuring accurate tracking without concurrent access issues.

4. Extensions

   Extension 1: HTTP POST: The application supports switching to a POST request by updating the SendCount function to post JSON data containing the unique request count.
   Extension 2: Load-Balanced Deduplication: For tracking unique IDs across multiple instances, a distributed or centralized data store, such as Redis, could replace sync.Map to maintain global uniqueness across instances.
   Extension 3: Distributed Streaming: For distributed logging, the application can send unique counts to a streaming service (e.g., Kafka) instead of writing to a file. This change supports distributed processing and scalability.

5. Logging and Monitoring

   Structured Logging: The application logs key information, including unique request counts and HTTP response statuses.
   Containerization: A Dockerfile is provided for easy deployment, ensuring consistent and scalable environments across development, testing, and production.

Conclusion

This application is optimized for high-throughput, stateless processing with extensions for load-balancing and scalability. The design leverages Go’s concurrency model and minimal dependencies, making it suitable for environments requiring high performance, efficient logging, and flexible integrations.
