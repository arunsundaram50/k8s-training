import fastapi, os, uvicorn, nest_asyncio, random

app = fastapi.FastAPI()
PORT = os.environ.get('PORT', '8080')
random_number = random.randint(1, 100)

@app.get("/upper/{text}")
def upper(text: str):
  
  with open('/data/upper.txt', 'at') as fp:
    print(f'converting {text} to upper', file=fp)

  return {
    "text": text.upper()
  }

nest_asyncio.apply()
uvicorn.run("upper:app", host="0.0.0.0", port=PORT, log_level="error")

