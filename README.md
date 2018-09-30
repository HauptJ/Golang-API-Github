# Golang-API-GitHub
Golang API that returns Github user / stargazer usernames and repositories as well as follower repositories.

### To Build and Run:
- **NOTE 1:** Port `8880` or the one specified in `docker-compose.yml` and `main.go` must be available.
- **NOTE 2:** [On Windows 10 Pro or greater? You can use Chocolatey to install everything.](https://chocolatey.org/)

1. Install Golang
  - [Golang Download and installation instructions](https://golang.org/dl/)


2. Install Golang dependencies
  - `go get github.com/gorilla/context`
  - `go get github.com/gorilla/mux`
  - `go get github.com/google/go-github/github`
  - `go get golang.org/x/oauth2`


3. **Optional:** Install Docker and Docker Compose
  - [Docker](https://docs.docker.com/install/)
  - [Docker Compose](https://docs.docker.com/compose/install/)


4. [Get a GitHub API key / token with the following permissions](https://github.com/settings/tokens/new)
  ![Github Token Permissions](GHPAT.png)

5. Set the token as an environment variable named `TOKEN`
  - **Windows:** `$env:TOKEN="TOKEN"`
  - **Linux:** `TOKEN="TOKEN"`


6. To run on host system
  1. `go build main.go`
  2. `./main` or `.\main.exe`


7. To run in a docker container
  - `docker-compose up`

### Endpoints:
#### /followers
- **TYPE:** GET
- Returns a list of 5 followers for the specified GithUb user as well as the followers of the followers 3 levels deep.
- `/followers/{:GitHub User}`
- **Example:** `/followers/torvalds`
- ![Followers Screenshot](GHFOLLOWERS.png)

#### /repos
- **TYPE:** GET
- Returns a list of 5 repositories for a specified GitHub user and a list of 5 stargazers for each repository. The lists are nested 3 levels deep.
- `/repos/{:GitHub User}`
- **Example:** `/repos/torvalds`
- ![Repos Screenshot](GHREPOS.png)
