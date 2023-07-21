FROM python:3.9

# RUN apt-get update && apt-get install -y python3-pip

# additional configuration
EXPOSE 8001

# update the image with needed files
COPY main.py main.py
COPY requirements.txt requirements.txt

# download the libraries into the image
RUN pip install --no-cache-dir -r requirements.txt

# run the application
CMD uvicorn main:app --reload --port 8001
