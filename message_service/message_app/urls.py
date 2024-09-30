from django.urls import path
from rest_framework.urlpatterns import format_suffix_patterns
from message_app import views

# Define what URLs are available in the API and how they map to views
urlpatterns = format_suffix_patterns([
    path('',                     views.api_root,                name='root'),
    path('messages/bulk-delete', views.api_delete_messages,     name='bulk-delete'),
    path('messages/fetch-new',   views.api_fetch_new_messages,  name='fetch-new-messages'),
    path('messages/',            views.MessageList.as_view(),   name='message-list'),
    path('messages/<int:pk>/',   views.MessageDetail.as_view(), name='message-detail'),
])