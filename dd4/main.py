import fastapi, os

app = fastapi.FastAPI()
access_count = 0
if os.path.exists("/data/count.txt"):
  with open("/data/count.txt", "rt") as fp:
    s = fp.read()
    access_count = int(s)

@app.get("/hello2")
def say_hello():
  global access_count
  access_count = access_count + 1
  with open("/data/count.txt", "wt") as fp:
    fp.write(str(access_count))
  return {
    "message": f"Hello, world, for the {access_count}'th time"
  }
