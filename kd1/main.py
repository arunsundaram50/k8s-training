import fastapi, os, uvicorn, nest_asyncio, random

app = fastapi.FastAPI()
PORT = os.environ.get('PORT', '8001')


@app.get("/hello")
def say_hello():
  random_number = random.randint(1, 100)
  return {
    "message": f"Hola, world {random_number}"
  }

nest_asyncio.apply()
uvicorn.run("main:app", host="0.0.0.0", port=PORT, log_level="error")

