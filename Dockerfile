FROM python:3.12

RUN mkdir /app

WORKDIR /app

ADD requirements.txt .

RUN pip install -r requirements.txt

ADD . .

CMD mkdocs serve -a 0.0.0.0:8000
