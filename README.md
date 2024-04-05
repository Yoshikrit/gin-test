# gin-test service
First gin web framework project practice

---

## Running the Application

To run the application locally, you need to have Go installed on your machine.

##### 1. Clone the repository:

```bash
git clone https://github.com/Yoshikrit/gin-test_go_project.git
cd <project-directory>

if you run in dev, you need to create docker-compose.yml for its own for postgresql database
then run on main

```bash
$env:APP_ENV = "development"
go run main.go
```

##### 2. Test project:

```bash
go test  ./...                     
go test gin-test/tests -v 
go test gin-test/handlers -coverage
```
but if you run in prod you need to build it as a container 

##### 3. build docker image:

```bash
docker build -t gin-test .
```

##### 4. run docker-compose.yml

```bash
docker-compose -f gintest-compose.yml up -d
```
And to use it, you need to connect to pgadmin and connect to server before using service in container