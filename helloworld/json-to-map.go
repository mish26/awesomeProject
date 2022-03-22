package main

import (
	"encoding/json"
	"fmt"
)

func dumpMap(space string, m map[string]interface{}) {
	for k, v := range m {
		if mv, ok := v.(map[string]interface{}); ok {
			fmt.Printf("{ \"%v\": \n", k)
			dumpMap(space+"\t", mv)
			fmt.Printf("}\n")
		} else {
			fmt.Printf("%v %v : %v\n", space, k, v)
		}
	}
}

var jsonStr = `
{
  "apparelFreeUsedItem": {
    "id": 99510001,
    "localizedName": "luxury_used_localized_name",
    "brand": {
      "id": "luxury-used-free-listing",
      "name": "LUXURY-USED-FREE-LISTING",
      "localizedName": "ラグジュアリ中古フリー出品",
      "description": "",
      "metaDescription": "",
      "title": "",
      "headDescription": "",
      "headContent": "",
      "mobileHeadContent": "",
      "tailContent": "",
      "mobileTailContent": "",
      "apparelSellCaution": "",
      "hideArticleRanking": false
    },
    "apparelCategory": {
      "id": 7,
      "name": "bag",
      "localizedName": "バッグ",
      "priority": 0
    },
    "apparelSubCategory": {
      "id": 1,
      "parentCategory": {
        "id": 0,
        "name": "",
        "localizedName": "",
        "priority": 0
      },
      "name": "hand",
      "localizedName": "ハンドバッグ",
      "priority": 0
    },
    "status": 0,
    "price": 14000,
    "commaPrice": "¥14,000",
    "sellerId": 3,
    "primaryPhoto": {
      "id": 0,
      "imageUrl": "https://cdn.dev.snkrdunk.com/中古メイン画像"
    },
    "photos": {
      "tops_back": {
        "id": 11,
        "imageUrl": "https://cdn.dev.snkrdunk.com/中古側面画像②"
      },
      "tops_collar": {
        "id": 12,
        "imageUrl": "https://cdn.dev.snkrdunk.com/中古上部画像"
      },
      "tops_front": {
        "id": 10,
        "imageUrl": "https://cdn.dev.snkrdunk.com/中古側面画像①"
      }
    },
    "optionPhotos": [],
    "condition": "newOldStock",
    "displayCondition": "新古品",
    "accessoriesNote": "accessories",
    "note": "詳細",
    "isDisplaySold": false,
    "version": 1,
    "apparelFreeUsedItemId": 0,
    "isLinkedToApparel": false,
    "apparelUsedItemPath" : "",
    "createdAt": "2018-05-10T15:00:00Z",
    "updatedAt": "2018-05-10T15:00:00Z"
  }
}
`

func main() {
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &jsonMap)
	if err != nil {
		panic(err)
	}
	dumpMap("", jsonMap)
}
