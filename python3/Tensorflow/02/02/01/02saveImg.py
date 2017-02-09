import requests
r = requests.get("http://uta.pw/shodou/img/3/3.png")

with open("test.png", "wb") as f:
    f.write(r.content)
print("saved")