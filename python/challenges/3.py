import urllib, re, os.path
url = "http://www.pythonchallenge.com/pc/def/equality.html"
fh = urllib.urlopen(url)
in_comment = False
chars = ""
for line in fh:
  if in_comment:
    chars += "".join(re.findall('[^A-Z][A-Z]{3}([a-z])[A-Z]{3}[^A-Z]',line))
  elif re.match("<!--",line):
    in_comment = True
  else:
    pass
print chars
