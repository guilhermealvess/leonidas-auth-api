<div align="center" id="top"> 
  <img src="./.github/app.gif" alt="Auth Api Jwt" />

  &#xa0;

  <!-- <a href="https://authapijwt.netlify.app">Demo</a> -->
</div>

<h1 align="center">Auth Api Jwt</h1>

<p align="center">
  <img alt="Github top language" src="https://img.shields.io/github/languages/top/guilhermealvess/auth-api-jwt?color=56BEB8">

  <img alt="Github language count" src="https://img.shields.io/github/languages/count/guilhermealvess/auth-api-jwt?color=56BEB8">

  <img alt="Repository size" src="https://img.shields.io/github/repo-size/guilhermealvess/auth-api-jwt?color=56BEB8">

  <img alt="License" src="https://img.shields.io/github/license/guilhermealvess/auth-api-jwt?color=56BEB8">

  <!-- <img alt="Github issues" src="https://img.shields.io/github/issues/{{YOUR_GITHUB_USERNAME}}/auth-api-jwt?color=56BEB8" /> -->

  <!-- <img alt="Github forks" src="https://img.shields.io/github/forks/{{YOUR_GITHUB_USERNAME}}/auth-api-jwt?color=56BEB8" /> -->

  <img alt="Github stars" src="https://img.shields.io/github/stars/guilhermealvess/auth-api-jwt?color=56BEB8" />
</p>

<!-- Status -->

<h4 align="center"> 
	🚧  Auth Api Jwt 🚀 Under construction...  🚧
</h4> 

<hr>

<p align="center">
  <a href="#dart-about">About</a> &#xa0; | &#xa0; 
  <a href="#sparkles-features">Features</a> &#xa0; | &#xa0;
  <a href="#rocket-technologies">Technologies</a> &#xa0; | &#xa0;
  <a href="#white_check_mark-requirements">Requirements</a> &#xa0; | &#xa0;
  <a href="#checkered_flag-starting">Starting</a> &#xa0; | &#xa0;
  <a href="#memo-license">License</a> &#xa0; | &#xa0;
  <a href="https://github.com/{{YOUR_GITHUB_USERNAME}}" target="_blank">Author</a>
</p>

<br>

## About ##

This project consists of being an API for authentication with JWT for other applications to validate the identification of their users.
The software design of the API used is the Clean Architecture, which aims to provide a decoupling between layers of our software.
At first we used gRPC as a way of communicating with the API, but the ease of changing or adding and making the application also a Rest API for example is extremely easy.

<img src="img/clean-arch.jpg" alt="Clean Architecture">

<!-- ## :sparkles: Features ##

:heavy_check_mark: Feature 1;\
:heavy_check_mark: Feature 2;\
:heavy_check_mark: Feature 3; -->

## :rocket: Technologies ##

The following tools were used in this project:

- [Golang](https://go.dev/)
- [gRPC](https://grpc.io/)
- [JWT](https://jwt.io/)
- [MongoDB](https://www.mongodb.com/)
- [Redis](https://redis.io/)
- [Docker](https://www.docker.com/)
## :white_check_mark: Requirements ##

Before starting :checkered_flag:, you need to have [Git](https://git-scm.com) and [Go 1.18](https://go.dev/) installed.

## :checkered_flag: Starting ##

```bash
# Clone this project
$ git clone https://github.com/guilhermealvess/auth-api-jwt

# Access
$ cd auth-api-jwt

# Install dependencies
$ go mod

# Run the project
$ go run main.go

# The server will initialize in the <http://localhost:50052>
# gRPC Client
evans -r repl --port=50052

```
## :memo: License ##

This project is under license from MIT. For more details, see the [LICENSE](LICENSE.md) file.


Made with :heart: by <a href="https://github.com/guilhermealvess" target="_blank">Guilherme Alves</a>

&#xa0;

<a href="#top">Back to top</a>
