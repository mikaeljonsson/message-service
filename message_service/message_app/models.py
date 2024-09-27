from django.db import models

# Create your models here.

"""
class Message(models.Model):
    # id field is added by default to be a bigint primary key
    created = models.DateTimeField(auto_now_add=True)
    recipient = models.CharField(max_length=200, blank=True, default='')
    title = models.CharField(max_length=100, blank=True, default='')
    code = models.TextField()
    linenos = models.BooleanField(default=False)
    language = models.CharField(choices=LANGUAGE_CHOICES, default='python', max_length=100)
    style = models.CharField(choices=STYLE_CHOICES, default='friendly', max_length=100)

    class Meta:
        ordering = ['created']
"""