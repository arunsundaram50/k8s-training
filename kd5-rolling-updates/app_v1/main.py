import fastapi, os, uvicorn, nest_asyncio, random

app = fastapi.FastAPI()
PORT = os.environ.get('PORT', '8001')
random_number = random.randint(1, 100)

@app.get("/hello")
def say_hello():
  return {
    "message": f"Hello, world V#1 {random_number}"
  }

nest_asyncio.apply()
uvicorn.run("main:app", host="0.0.0.0", port=PORT, log_level="error")

