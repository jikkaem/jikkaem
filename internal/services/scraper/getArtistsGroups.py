import csv

ang = set()
with open("./kpop-idols.csv", 'r') as file:
  csvreader = csv.reader(file)
  for row in csvreader:
    ang.add(row[1])
    if row[5] != "":
        ang.add(row[5])

f = open("listArtists.txt", 'a')
for elem in ang:
  f.write(elem + ",")
f.close()
    
