import fastapi, os, uvicorn, nest_asyncio, random, socket, sys

app = fastapi.FastAPI()
PORT = os.environ.get('PORT', '8001')
SERVER_IP = socket.gethostbyname(socket.gethostname())

@app.get("/hello")
def say_hello():
  random_number = random.randint(1, 100)
  print(f"GET request /hello {SERVER_IP=}")
  return {
    "message": f"Hola, world {SERVER_IP=}: {random_number}"
  }

print(f"Listening to port ${PORT=}...", flush=True, file=sys.stderr)
nest_asyncio.apply()
uvicorn.run("main:app", host="0.0.0.0", port=PORT, log_level="error")

