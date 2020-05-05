import urllib.request
import collections
from bs4 import BeautifulSoup
import simplejson as json

categories = []

Category = collections.namedtuple('Category', ['id', 'name', 'sats'])

for i in range(2, 53):
    print(i)
    f = urllib.request.urlopen('https://www.n2yo.com/satellites/?c=' + str(i))
    data = f.read().decode('utf-8')
    soup = BeautifulSoup(data, "html.parser")
    title = soup.find('h1').text
    title = title.replace(' SATELLITES', '')
    sats = []
    tables = soup.find_all('table')
    for tr in tables[1].find_all('tr'):
        row = tr.find_all('td')
        if len(row) == 6:
            sat = int(int(row[1].text))
            sats.append(sat)

    categories.append(Category(i, title, sats))

with open('categories.json', 'w') as f:
    json.dump(categories, f, sort_keys=True, indent=4 * ' ')
