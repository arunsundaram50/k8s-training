import fastapi, os, uvicorn, nest_asyncio, requests, sys

app = fastapi.FastAPI()
PORT = os.environ.get('PORT', '8001')
HELLO_HOST = os.environ.get('hello_host')
HELLO_PORT = os.environ.get('hello_port')


@app.get("/upper/{text}")
def get_upper(text):
  URL = f"http://{HELLO_HOST}:{HELLO_PORT}/hello/{text}"
  print(f"{URL=}", flush=True, file=sys.stderr)
  resp = requests.get(URL)
  print(resp.json(), flush=True, file=sys.stderr)
  message = resp.json()["message"]
  return {
    "text": f"{message.upper()}"
  }

nest_asyncio.apply()
uvicorn.run("main:app", host="0.0.0.0", port=PORT, log_level="error")

