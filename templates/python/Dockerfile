FROM python:3-alpine

WORKDIR /usr/src/app

RUN pip install --no-cache-dir nose coverage flake8

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY . .
RUN pip install -e .
