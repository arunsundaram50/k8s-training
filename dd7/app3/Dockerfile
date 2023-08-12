FROM python:alpine

# RUN apt-get update && apt-get install -y python3-pip
EXPOSE 8080

# update the image with needed files
COPY greet_in_upper.py greet_in_upper.py
COPY requirements.txt requirements.txt

# download the libraries into the image
RUN pip install --no-cache-dir -r requirements.txt

# run the application
CMD python3 ./greet_in_upper.py


