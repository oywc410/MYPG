import urllib.request

url = "http://uta.pw/shodou/img/28/214.png"
savername = "test.png"

mem = urllib.request.urlopen(url).read()

with open(savername, mode="wb") as f:
    f.write(mem)
    print("保存しました")