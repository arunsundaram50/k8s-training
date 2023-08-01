FROM python:alpine


# Usage examples:
# ===============
#   docker image build -t dd5b
#   docker container run --rm dd5b
#   docker container run --rm -e op=upper dd5b
#   docker container run --rm -e op=upper dd5b world


# image build time
ENV op="lower"
COPY convert-case.py ./convert-case.py


# container run time
ENTRYPOINT ["./convert-case.py"]
CMD ["Hello"]

