from django.db import models

# Create your models here.
class Message(models.Model):
    # id field is added by default to be a bigint auto-incrementing primary key
    create_time = models.DateTimeField(auto_now_add=True)
    recipient = models.CharField(max_length=200)
    message_body = models.CharField(max_length=100, blank=True, default='') # empty message is allowed
    is_fetched = models.BooleanField(default=False)

    class Meta:
        # The requirement is to order in chrononogial order, but for efficiency
        # in a real time scenario, it would perhaps be more efficient to order them by
        # ['recipient', 'id'], as the search would then normally always be done per recipient
        # and the entire DB does not need to be searched through.
        ordering = ['id']
