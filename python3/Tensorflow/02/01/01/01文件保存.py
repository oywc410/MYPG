import urllib.request

url = "http://uta.pw/shodou/img/28/214.png"
savername = "test.png"

#保存文件
urllib.request.urlretrieve(url, savername)