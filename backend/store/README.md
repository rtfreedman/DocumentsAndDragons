# Store Description
## Character
Examples:
```json
{
	"_id" : ObjectId("5cce3c310a4836dbecd22fa4"),
	"level" : 1,
	"name" : "Rorik Ironforge",
	"Abilities" : {
		"Rage" : {
			"Charges" : 0,
			"MaxCharges" : 3,
			"FilterSelf" : {
				"filter" : {
					"_id" : "5cce3c310a4836dbecd22fa4",
				},
				"update" : {
					"$inc" : {
						"Abilities.Rage.Charges" : -1
					}
				}
			}
		}
	}
}
```