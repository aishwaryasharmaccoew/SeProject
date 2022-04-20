# findmyknife Backend

# Setup Guide

1. Run the following command to install the required go modules
```shell
cd backend
go mod download
go install
```

2. Run the backend server
```shell
cd src
go run .
```

# Start using Docker
```shell
cd backend
docker build --tag fmk_backend .
docker run -p 5001:5001 fmk_backend
```

