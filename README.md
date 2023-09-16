# go-url-shortener
This project focuses on the generation of short URLs and QR code images, as well as the redirection of short URLs to their original counterparts. It makes use of the Go programming language with the following libraries: go-gin, gorm, go-wire, and PostgreSQL.

You can access the service here: https://golang-url-shortener.onrender.com/

# Approach
The approach taken here is quite straightforward. It generates short URLs using a randomly generated 6-character string. To prevent data collisions, a unique key constraint is applied within the database. If a collision occurs, the system attempts to generate a new random string, with up to three retries before marking the operation as failed.

# Todo
There are several improvements and additional features that can be considered for this project:
1. Use Redis for Improved Performance: Implementing Redis can significantly enhance the performance of the URL shortening service by caching frequently accessed URLs.
2. Add Short URL Expiration: Implementing an expiration time for short URLs can enhance security and manageability. This ensures that short URLs have a limited lifespan.
3. Explore Advanced Short URL Generation: Investigate more robust methods for generating short URLs to reduce the risk of collisions further.
