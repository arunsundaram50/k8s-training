FROM python:alpine

# RUN apt-get update && apt-get install -y python3-pip

# update the image with needed files
COPY upper.py upper.py
COPY requirements.txt requirements.txt

# download the libraries into the image
RUN pip install --no-cache-dir -r requirements.txt

# run the application
CMD python3 ./upper.py


