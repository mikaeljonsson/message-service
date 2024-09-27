from django.urls import path
from message_app import views

urlpatterns = [
    path('messages/', views.message_list),
    path('messages/<int:pk>/', views.message_detail),
]