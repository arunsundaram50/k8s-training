FROM python:3.9

# additional configuration
EXPOSE 8001

# update the image with needed files
COPY main.py main.py
COPY requirements.txt requirements.txt


# download the libraries into the image
RUN pip install --no-cache-dir -r requirements.txt

# run the application
CMD uvicorn main:app --reload --host 0.0.0.0 --port 8001
