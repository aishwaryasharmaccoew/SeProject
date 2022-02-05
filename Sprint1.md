# Introduction
Creating an E-commerce website for Swiss Knives as part of Software Engineer course work.

# Team Members
1. Aishwarya Sharma
2. Nimish Bajaj
3. Siddhant Jain
4. Vaibhav Mishra

# Tech Stack
1. Angular
2. Go
3. Ionic
4. Rxjs
5. SQLite
6. Gin
7. Gorm

# Backend Feature Set achieved
## v0.1 
1. List items on the webpage
2. Filter for an item using one of the filters
3. Look at the product description page for an item


# FRONT-END  





# BACKEND APIs Implemented

**IMPORTANT:** Import the below **postman** collection and use the APIs
Backend API Documentation: [Link](https://github.com/aishwaryasharmaccoew/SeProject/blob/main/backend/api_documentation.md)  
Postman Collection (v2): [Link](https://github.com/aishwaryasharmaccoew/SeProject/blob/main/backend/src/postman_api_samples/findmyknife.postman_collection.json)

## Landing Page API / All products API
**Type** GET

**URL:** localhost:5001/

## Get Prodcut Details
**Type** GET

**URL:** localhost:5001/product/:id

**Sample Request:**
```
localhost:5001/product/B-0.2373.T
```


## Product Filter API
This API accepts an array of filter elements on the product tools and returns the filtered list of products.

**Type** POST

**URL:** localhost:5001/product 

**Sample Payload:**
```json
{
	"filters": ["screw", "scissor", "large"]
}
```


## API demo using command line and postman

