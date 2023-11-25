
# Docker Containerized Go WebAPI

This is a simple project created to showcase what I've learned while attending the DevOps Learning Path on KodeKloud.

## Contains

- Simple CRUD API using GoLang.
- Simple test cases for every endpoint.
- `Dockerfile` to containerize the WebAPI with Docker.
- `docker-compose.yml` file is designed to orchestrate the deployment of a Go WebAPI and a MySQL database for development or production environments. It leverages Docker Compose to define and manage the services, their dependencies, and configurations.
- `Jenkinsfile` A declarative pipeline that performs the following steps:
    1. Cloning this git repository.
    2. Running tests using the `go` command.
    3. Building Go WebAPI into an `executable` file.
    4. Building an image using Docker.
    5. Pushing the image into a private repository.

## Requirements for Jenkins
- [Go v1.21.1](https://go.dev/doc/install) installed on the Jenkins instance.
- The following plugins should be installed on the Jenkins instance:
    - [Git Jenkins Plugin](https://plugins.jenkins.io/git/)
    - [Docker Pipeline Plugin](https://plugins.jenkins.io/docker-workflow/)
- A private Docker Hub repository
    <br>(change the `repo` value inside the Jenkinsfile).
- Create a credential entry for Docker with the ID `docker-credentials`.
- Create a new Jenkins pipeline job that pulls the definition from Git SCM.

## Environment Variables

To run this project, you will need to create a `.env` file with the variables provided in the `.env.example` file.

## API Reference

#### Get all products

```http
  GET /products
```

#### Get product

```http
  GET /products/${id}
```

| Parameter | Type     | Description                       | Required
| :-------- | :------- | :-------------------------------- | :--------
| `id`      | `string` | Id of product to fetch | Yes

#### Create product

```http
  POST /products
```
Attributes to pass in JSON format within the request:

| Attribute  | Type     | Required |
| :--------- | :------- | :------- |
| `name`     | `string` | Yes      |
| `price`    | `float`  | Yes      |
| `quantity` | `int`    | Yes      |


#### Update product
```http
  PUT /products/${id}
```

| Parameter | Type     | Description                       | Required
| :-------- | :------- | :-------------------------------- | :--------
| `id`      | `string` | Id of product to update | Yes

Attributes to pass in JSON format within the request:

| Attribute  | Type     | Required |
| :--------- | :------- | :------- |
| `name`     | `string` | Yes      |
| `price`    | `float`  | Yes      |
| `quantity` | `int`    | Yes      |

#### Delete product

```http
  DELETE /products/${id}
```

| Parameter | Type     | Description                       | Required
| :-------- | :------- | :-------------------------------- | :--------
| `id`      | `string` | Id of product to delete | Yes





