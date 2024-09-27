from django.db import models

# Create your models here.
class Message(models.Model):
    # id field is added by default to be a bigint auto-incrementing primary key
    create_time = models.DateTimeField(auto_now_add=True)
    recipient = models.CharField(max_length=200)
    message_body = models.CharField(max_length=100, blank=True, default='') # empty message is allowed
    is_fetched = models.BooleanField(default=False)

    class Meta:
        ordering = ['recipient','id']
