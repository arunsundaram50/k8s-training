import fastapi, os, uvicorn, nest_asyncio, random

app = fastapi.FastAPI()
PORT = os.environ.get('PORT', '8082')
random_number = random.randint(1, 100)

@app.get("/upper/{text}")
def upper(text: str):
  return {
    "text": text.upper()
  }

nest_asyncio.apply()
uvicorn.run("upper:app", host="0.0.0.0", port=PORT, log_level="error")

