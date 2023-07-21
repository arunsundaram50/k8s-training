import fastapi

app = fastapi.FastAPI()


@app.get("/hello")
def say_hello():
  return {
    "message": "Hello, world"
  }


