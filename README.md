# INTRODUCTION
1. A simple URL shortenerservice that will accept a URL as an argument over a REST API and
return a shortened URL as a result. Please visit https://bitly.com/ as an example.
2. The application has 3 API,
   * API 1 - Shortens the url
   * API 2 - Redirects the shortened url to original address
   * API 3 - A metrics API which gives the Top 3 hit domains.
For eg. if the user hasshortened 4 YouTube video links, 1 StackOverflow link, 2
Wikipedia links and 6 Udemy tutorial links. Then the output would be:
Udemy: 6
YouTube: 4
Wikipedia: 2
3. Dockerfile for this application which builds a successful docker image.

# TECH-STACK
* go version - 1.22.0
* framework - Gorilla mux
* IDE : Intellij IDEA
* Http Client : Resty
* Postman
* Cache - gocache
* Docker Desktop - 4.28.0

# FAQ
* Total APIs - 3
* use Makefile to build and run the project

# IMAGES
1. Shortening URL : /url-shortener/shorten - POST
   ![API-1.png](images/API-1.png)
2. Redirecting URL : /redirect - GET
   ![API-2.png](images/API-2.png)
3. Metrics API : /url-shortener/metrics - GET
   ![API-3.png](images/API-3.png)
4. Docker Image built successfully
   ![Docker-1.png](images/Docker-1.png)
5. Docker image running successfully in Docker desktop
   ![Docker-2.png](images/Docker-2.png)
6. executable formed in Docker
   ![Docker-3.png](images/Docker-3.png)
