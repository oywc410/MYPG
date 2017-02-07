from bs4 import BeautifulSoup

html = """
<html><body>
<h1 id="title">h1 text</h1>
<p id="body">body text</p>
</body></html>
"""

soup = BeautifulSoup(html, 'html.parser')

title = soup.find(id="title")
body = soup.find(id="body")

print("#title=" + title.string)
print("#body=" + body.string)