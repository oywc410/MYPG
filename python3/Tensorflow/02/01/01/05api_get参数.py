import sys
import urllib.request as req
import urllib.parse as parse

if len(sys.argv) <= 1:
    print("USACE: byakunnin.py (keyword)")
    sys.exit()

keyword = sys.argv[1]

API = "http://api.aoikujira.com/hyakunin/get.php"
query = {
    "fmt": "ini",
    "key": keyword,
}

params = parse.urlencode(query)
url = API + "?" + params
print("url=", url)

with req.urlopen(url) as r:
    b = r.read()
    data = b.decode('utf-8')
    print(data)