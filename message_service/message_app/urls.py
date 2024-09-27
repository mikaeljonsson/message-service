from django.urls import path
from rest_framework.urlpatterns import format_suffix_patterns
from message_app import views

urlpatterns = [
    path('messages/', views.message_list),
    path('messages/<int:pk>/', views.message_detail),
]

urlpatterns = format_suffix_patterns(urlpatterns)
