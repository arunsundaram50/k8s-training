import fastapi, os, uvicorn, nest_asyncio, random

app = fastapi.FastAPI()
PORT = os.environ.get('PORT', '8081')

@app.get("/greet/{user}")
def greet(user: str):
  random_number = random.randint(1, 100)
  return {
    "message": f"Hello, {random_number}: {user}"
  }

nest_asyncio.apply()
uvicorn.run("greet:app", host="0.0.0.0", port=PORT, log_level="error")

