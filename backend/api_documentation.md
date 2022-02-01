# APIs

**IMPORTANT:** Import the **postman** collection and use the APIs

## Filter API
This API accepts an array of filter elements on the product tools and returns the filtered list of products.

**URL:**  
localhost:5001/filter 

**Sample Payload:**
```json
{
	"filters": ["screw", "scissor", "large"]
}
```

