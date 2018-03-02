FROM python:3-alpine

# Os dependencies
# RUN apk add --update --no-cache \

WORKDIR /usr/src/{{.Name}}

# Python dev dependencies
RUN pip install --no-cache-dir nose coverage flake8

# Python production dependencies
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
