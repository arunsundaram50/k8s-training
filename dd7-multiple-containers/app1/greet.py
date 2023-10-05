import fastapi, os, uvicorn, nest_asyncio, random

app = fastapi.FastAPI()
PORT = os.environ.get('PORT', '8080')

@app.get("/greet/{user}")
def greet(user: str):
  print(f'processing request for {user}')
  random_number = random.randint(1, 100)
  return {
    "message": f"Hello, {user} ({random_number})"
  }

print(f'Listening to port {PORT}...')
nest_asyncio.apply()
uvicorn.run("greet:app", host="0.0.0.0", port=PORT, log_level="error")

