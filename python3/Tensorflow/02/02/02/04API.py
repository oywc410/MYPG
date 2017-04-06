import requests
import json

apiKey = "xxxxxxxxxxxx"

cities = ["Tokyo,JP", "London,UK", "New York,US"]
api = "http://xxxxxxxxxxxxxxx/?p={city}&APPID={key}"

for name in cities :
    url = api.format(city=name, key=apiKey)
    r = requests.get(url)
    data = json.loads(r.text)
    data["name"]
