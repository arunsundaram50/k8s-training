import fastapi, os, uvicorn, nest_asyncio, random, socket

app = fastapi.FastAPI()
PORT = os.environ.get('PORT', '8001')
SERVER_IP = socket.gethostbyname(socket.gethostname())

@app.get("/hello")
def say_hello():
  random_number = random.randint(1, 100)
  return {
    "message": f"Hola, world {SERVER_IP=}: {random_number}"
  }

nest_asyncio.apply()
uvicorn.run("main:app", host="0.0.0.0", port=PORT, log_level="error")

