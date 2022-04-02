# APIs

**IMPORTANT:** Import the **postman** collection and use the APIs

## Landing Page API / All products API
**Type** GET

**URL:** localhost:5001?numResults=10&pageNum=2

**Sample Request:**
```
localhost:5001/product?numResults=10&pageNum=1
```



## Get Prodcut Details
**Type** GET

**URL:** localhost:5001/product/:id

**Sample Request:**
```
localhost:5001/product/B-0.2373.T
```
**Sample Response:**

```json
[
    {
        "Id": "B-0.7100.T",
        "Name": "Swiss Card Classic",
        "Price": 33,
        "Category": "SAK / Swiss Army Knives",
        "ImageUrl": "https://imageengine.victorinox.com/mediahub/33093/420Wx368H/SAK_0_7100_T__S1.jpg",
        "Product_url": "https://www.victorinox.com/global/en/Products/Swiss-Army-Knives/Swiss-Cards/Swiss-Card-Classic/p/0.7100.T",
        "Colors": "black transparent,blue transparent,red transparent",
        "Key_features": "",
        "Tools": "nail file,screwdriver 2.5 mm,toothpick,tweezers,pressurized ballpoint pen"
    },
    {
        "Id": "B-0.7300.T",
        "Name": "Swiss Card Lite",
        "Price": 39,
        "Category": "SAK / Swiss Army Knives",
        "ImageUrl": "https://imageengine.victorinox.com/mediahub/33145/420Wx368H/SAK_0_7300_T__S1.jpg",
        "Product_url": "https://www.victorinox.com/global/en/Products/Swiss-Army-Knives/Swiss-Cards/Swiss-Card-Lite/p/0.7300.T",
        "Colors": "black transparent,blue transparent,red transparent",
        "Key_features": "",
        "Tools": "scissors,magnifying glass,screwdriver 3 mm,tweezers,pressurized ballpoint pen"
    },
    {
        "Id": "B-0.8201.26",
        "Name": "Pioneer Alox",
        "Price": 39,
        "Category": "SAK / Swiss Army Knives",
        "ImageUrl": "https://imageengine.victorinox.com/mediahub/40425/420Wx368H/SAK_0_8201_26__S1.jpg",
        "Product_url": "https://www.victorinox.com/global/en/Products/Swiss-Army-Knives/Medium-Pocket-Knives/Pioneer-Alox/p/0.8201.26",
        "Colors": "silver",
        "Key_features": "",
        "Tools": "large blade,reamer, punch,can opener,screwdriver 3 mm"
    },
    {
        "Id": "B-0.8241.26",
        "Name": "Farmer Alox",
        "Price": 45,
        "Category": "SAK / Swiss Army Knives",
        "ImageUrl": "https://imageengine.victorinox.com/mediahub/33243/420Wx368H/SAK_0_8241_26__S1.jpg",
        "Product_url": "https://www.victorinox.com/global/en/Products/Swiss-Army-Knives/Medium-Pocket-Knives/Farmer-Alox/p/0.8241.26",
        "Colors": "silver",
        "Key_features": "",
        "Tools": "large blade,reamer, punch,can opener,screwdriver 3 mm,bottle opener"
    },
    {
        "Id": "B-0.8341.MC9",
        "Name": "Hunter XT Grip",
        "Price": 59,
        "Category": "SAK / Swiss Army Knives",
        "ImageUrl": "https://imageengine.victorinox.com/mediahub/33259/420Wx368H/SAK_0_8341_MC9__S1.jpg",
        "Product_url": "https://www.victorinox.com/global/en/Products/Swiss-Army-Knives/Large-Pocket-Knives/Hunter-XT-Grip/p/0.8341.MC9",
        "Colors": "orange/black",
        "Key_features": "",
        "Tools": "large blade,gutting blade,wood saw"
    },
    {
        "Id": "B-0.8413.3",
        "Name": "Sentinel",
        "Price": 29,
        "Category": "SAK / Swiss Army Knives",
        "ImageUrl": "https://imageengine.victorinox.com/mediahub/33290/420Wx368H/SAK_0_8413_3__S1.jpg",
        "Product_url": "https://www.victorinox.com/global/en/Products/Swiss-Army-Knives/Large-Pocket-Knives/Sentinel/p/0.8413.3",
        "Colors": "black",
        "Key_features": "",
        "Tools": "large blade,tweezers,toothpick,key ring"
    },
    {
        "Id": "B-0.8413.M3",
        "Name": "Sentinel One Hand",
        "Price": 29,
        "Category": "SAK / Swiss Army Knives",
        "ImageUrl": "https://imageengine.victorinox.com/mediahub/33294/420Wx368H/SAK_0_8413_M3__S1.jpg",
        "Product_url": "https://www.victorinox.com/global/en/Products/Swiss-Army-Knives/Large-Pocket-Knives/Sentinel-One-Hand/p/0.8413.M3",
        "Colors": "black",
        "Key_features": "",
        "Tools": "large blade,tweezers,toothpick,key ring"
    },
    {
        "Id": "B-0.8461.MWCH",
        "Name": "Swiss Soldier's Knife 08",
        "Price": 51,
        "Category": "SAK / Swiss Army Knives",
        "ImageUrl": "https://imageengine.victorinox.com/mediahub/33311/420Wx368H/SAK_0_8461_MWCH__S1.jpg",
        "Product_url": "https://www.victorinox.com/global/en/Products/Swiss-Army-Knives/Large-Pocket-Knives/Swiss-Soldier%27s-Knife-08/p/0.8461.MWCH",
        "Colors": "green/black",
        "Key_features": "",
        "Tools": "large blade with wavy edge,reamer, punch,bottle opener, lockable,wire stripper,screwdriver 7.5 mm"
    },
    {
        "Id": "B-0.8463.3",
        "Name": "Trailmaster",
        "Price": 49,
        "Category": "SAK / Swiss Army Knives",
        "ImageUrl": "https://imageengine.victorinox.com/mediahub/33323/420Wx368H/SAK_0_8463__S1.jpg",
        "Product_url": "https://www.victorinox.com/global/en/Products/Swiss-Army-Knives/Large-Pocket-Knives/Trailmaster/p/0.8463",
        "Colors": "Camouflage,red",
        "Key_features": "",
        "Tools": "large blade,bottle opener, lockable,wire stripper,screwdriver 7.5 mm,wood saw,can opener"
    },
    {
        "Id": "B-0.8623.MWN",
        "Name": "Rescue Tool",
        "Price": 99,
        "Category": "SAK / Swiss Army Knives",
        "ImageUrl": "https://imageengine.victorinox.com/mediahub/33373/420Wx368H/SAK_0_8623_MWN__S1.jpg",
        "Product_url": "https://www.victorinox.com/global/en/Products/Swiss-Army-Knives/Large-Pocket-Knives/Rescue-Tool/p/0.8623.MWN",
        "Colors": "phosphorescent yellow",
        "Key_features": "",
        "Tools": "large blade with wavy edge,seatbelt cutter,reamer, punch,disc saw for shatterproof glass,bottle opener, lockable,wire stripper,screwdriver 7.5 mm"
    }
]
```

## Product Filter API
This API accepts an array of filter elements on the product tools and returns the filtered list of products.

**Type** POST

**URL:** localhost:5001/product?numResults=10&pageNum=1

**Sample Request:**
```
localhost:5001/product?numResults=10&pageNum=1
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
    ]
}
```

