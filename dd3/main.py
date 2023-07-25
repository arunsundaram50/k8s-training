import fastapi

app = fastapi.FastAPI()
access_count = 0

@app.get("/hello2")
def say_hello():
  global access_count
  access_count = access_count + 1
  return {
    "message": f"Hello, world, for the {access_count}'th time"
  }


