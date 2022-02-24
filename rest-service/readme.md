## rest-service

### any extras you can throw in?

Logging for insight for backend and test script for the client side.

### How would you test this service?

Postman / curl endpoints for status and response. Unit test for the business
logic in the package.

### How would you audit it?

Logs

### How would an ops person audit it?

Logs. Send logs to Loki or ELK. OR store in file and rotate.  

### monitor

Prometheus blackbox for endpoints
