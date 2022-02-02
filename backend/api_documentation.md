# APIs

**IMPORTANT:** Import the **postman** collection and use the APIs

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

