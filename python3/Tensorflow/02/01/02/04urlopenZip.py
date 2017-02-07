from bs4 import BeautifulSoup
import urllib.request as req

url = "http://api.aoikujira.com/zip/xml/1500042"

res = req.urlopen(url)

soup = BeautifulSoup(res, "html.parser")

ken = soup.find("ken").string
shi = soup.find("shi").string
cho = soup.find("cho").string

print(ken, shi, cho)