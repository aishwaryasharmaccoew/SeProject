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

**URL:** localhost:5001/product/:id

**Sample Request:**
```
localhost:5001/product/B-0.2373.T
```

**Sample Payload:**
```json
{
	"filters": ["screw", "scissor", "large"]
}
```
**Sample Response:**
```json
{
    "items": [
        {
            "Id": "B-0.6203",
            "Name": "Classic",
            "Price": 18,
            "Category": "SAK / Swiss Army Knives",
            "ImageUrl": "https://imageengine.victorinox.com/mediahub/32624/420Wx368H/SAK_0_6203__S1.jpg",
            "Product_url": "https://www.victorinox.com/global/en/Products/Swiss-Army-Knives/Small-Pocket-Knives/Classic/p/0.6203",
            "Colors": "red",
            "Key_features": "",
            "Tools": "small blade,scissors,nail file,nail cleaner"
        },
        {
            "Id": "B-0.6221.26",
            "Name": "Classic Alox",
            "Price": 27,
            "Category": "SAK / Swiss Army Knives",
            "ImageUrl": "https://imageengine.victorinox.com/mediahub/32641/420Wx368H/SAK_0_6221_26__S1.jpg",
            "Product_url": "https://www.victorinox.com/global/en/Products/Swiss-Army-Knives/Small-Pocket-Knives/Classic-Alox/p/0.6221.26",
            "Colors": "silver",
            "Key_features": "",
            "Tools": "small blade,scissors,nail file,screwdriver 2.5 mm,key ring"
        },
        {
            "Id": "B-0.6223",
            "Name": "Classic SD Printed",
            "Price": 23,
            "Category": "SAK / Swiss Army Knives",
            "ImageUrl": "https://imageengine.victorinox.com/mediahub/35912/420Wx368H/SAK_0_6223_942__S1.jpg",
            "Product_url": "https://www.victorinox.com/global/en/Products/Swiss-Army-Knives/Small-Pocket-Knives/Classic-SD-Printed/p/0.6223.942",
            "Colors": "Navy Camouflage,Desert Camouflage,Camouflage,Chocolate,VX Colors",
            "Key_features": "",
            "Tools": "small blade,scissors,nail file,screwdriver 2.5 mm"
        },
        {
            "Id": "B-0.6225",
            "Name": "Signature",
            "Price": 23,
            "Category": "SAK / Swiss Army Knives",
            "ImageUrl": "https://imageengine.victorinox.com/mediahub/32932/420Wx368H/SAK_0_6225__S1.jpg",
            "Product_url": "https://www.victorinox.com/global/en/Products/Swiss-Army-Knives/Small-Pocket-Knives/Signature/p/0.6225",
            "Colors": "blue transparent,red transparent,red",
            "Key_features": "",
            "Tools": "toothpick,small blade,scissors,nail file"
        },
        {
            "Id": "B-0.6226",
            "Name": "Signature Lite",
            "Price": 37,
            "Category": "SAK / Swiss Army Knives",
            "ImageUrl": "https://imageengine.victorinox.com/mediahub/32954/420Wx368H/SAK_0_6226__S1.jpg",
            "Product_url": "https://www.victorinox.com/global/en/Products/Swiss-Army-Knives/Small-Pocket-Knives/Signature-Lite/p/0.6226",
            "Colors": "SilverTech,blue transparent,red transparent,red",
            "Key_features": "",
            "Tools": "small blade,scissors,nail file,screwdriver 2.5 mm"
        }
    ],
    "next_page_id": "B-0.6540.16"
}
```

