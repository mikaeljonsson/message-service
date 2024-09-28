from django.urls import path
from rest_framework.urlpatterns import format_suffix_patterns
from message_app import views

urlpatterns = format_suffix_patterns([
    path('', views.api_root),
    path('messages/bulk-delete', views.api_delete_messages),
    path('messages/',
        views.MessageList.as_view(),
        name='message-list'),
    path('messages/<int:pk>/',
        views.MessageDetail.as_view(),
        name='message-detail'),

    # path('users/',
    #    views.UserList.as_view(),
    #    name='user-list'),
    #path('users/<int:pk>/',
    #    views.UserDetail.as_view(),
    #    name='user-detail')
])