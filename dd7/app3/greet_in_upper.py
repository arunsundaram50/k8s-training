import fastapi, os, uvicorn, nest_asyncio, random, requests

app = fastapi.FastAPI()
PORT = os.environ.get('PORT', '8080')
random_number = random.randint(1, 100)

@app.get("/greet-in-upper/{text}")
def greet_in_upper(text: str):
  msg = requests.get(f'http://app1:8080/greet/{text}').json()['message']
  upper_greet = requests.get(f'http://app2:8080/upper/{msg}').json()['text']
  return {
    "upper-greet": upper_greet
  }

if __name__ == "__main__":
  nest_asyncio.apply()
  uvicorn.run("greet_in_upper:app", host="0.0.0.0", port=int(PORT), log_level="error", reload = True)

