FROM python:3.9-slim
WORKDIR /app
COPY . .
# COPY message_service/requirements.txt .
WORKDIR /app/message_service/
RUN pip install --no-cache-dir -r requirements.txt

EXPOSE 8000
RUN python manage.py makemigrations message_app
RUN python manage.py migrate message_app

CMD ["python", "manage.py", "runserver", "0.0.0.0:8000"]
