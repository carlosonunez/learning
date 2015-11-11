import urllib, re, os.path
url = "http://www.pythonchallenge.com/pc/def/ocr.html"
fh = urllib.urlopen(url)
in_comment = False
rare_chars = ""
for line in fh:
  if in_comment:
    letters_found = re.findall('[A-Za-z]',line)
    rare_chars += "".join(re.findall('[A-Za-z]',line)
  elif re.match("find rare characters in the mess below",line):
    in_comment = True
  else:
    pass
print rare_chars
