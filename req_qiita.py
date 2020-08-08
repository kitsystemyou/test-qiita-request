import requests
import pprint

with open('./q_token.txt') as f:
    qiita_access_token = f.read().strip()

header = {'Authorization': 'Bearer {}'.format(qiita_access_token)}
print(header)
print("conflict test2")
