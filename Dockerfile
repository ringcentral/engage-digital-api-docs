FROM python:3

RUN mkdir /app

WORKDIR /app

ADD . .

RUN pip install -r requirements.txt

CMD mkdocs serve -a 0.0.0.0:8000
