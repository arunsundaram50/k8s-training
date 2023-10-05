import fastapi, os, uvicorn, nest_asyncio, sys, random

app = fastapi.FastAPI()
access_count = 0
PORT = 8001
INSTANCE_ID = random.randint(1, 100)

@app.get("/hello/{user}")
def say_hello(user):
  global access_count
  access_count = access_count + 1
  GREET_PREFIX = os.environ.get('GREET_PREFIX', 'Hello')
  print(f'GET /hello for {INSTANCE_ID=}, {user=}, {access_count}th time', file=sys.stderr, flush=True)
  return {
    "message": f"{INSTANCE_ID=}, {GREET_PREFIX}, {user}, for the {access_count}'th time"
  }

print(f'Listening to {PORT=}...', file=sys.stderr, flush=True)
nest_asyncio.apply()
uvicorn.run("main:app", host="0.0.0.0", port=PORT, log_level="error")
