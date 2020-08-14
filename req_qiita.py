import requests
import os

qiita_access_token = os.getenv("Q_TOKEN")

header = {'Authorization': 'Bearer {}'.format(qiita_access_token)}
print(header)

url_items = 'https://qiita.com/api/v2/authenticated_user/items'
query_params = {
    "page":1,
    "per_page":20
}

response = requests.get(url=url_items, headers =header, params=query_params)
print(response.status_code, type(response.json))
json = response.json()
print(json[:][0]["title"])
print([x["title"] for x in json])
