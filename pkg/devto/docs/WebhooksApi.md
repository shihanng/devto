# \WebhooksApi

All URIs are relative to *https://dev.to/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhook**](WebhooksApi.md#CreateWebhook) | **Post** /webhooks | Create a new webhook
[**DeleteWebhook**](WebhooksApi.md#DeleteWebhook) | **Delete** /webhooks/{id} | A webhook endpoint
[**GetWebhookById**](WebhooksApi.md#GetWebhookById) | **Get** /webhooks/{id} | A webhook endpoint
[**GetWebhooks**](WebhooksApi.md#GetWebhooks) | **Get** /webhooks | Webhooks



## CreateWebhook

> WebhookShow CreateWebhook(ctx, optional)

Create a new webhook

This endpoint allows the client to create a new webhook.  \"Webhooks\" are used to register HTTP endpoints that will be called once a relevant event is triggered inside the web application, events like `article_created`, `article_updated`. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateWebhookOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookCreate** | [**optional.Interface of WebhookCreate**](WebhookCreate.md)| Webhook to create | 

### Return type

[**WebhookShow**](WebhookShow.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhook

> DeleteWebhook(ctx, id)

A webhook endpoint

This endpoint allows the client to delete a single webhook given its `id`. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| Id of the webhook | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookById

> WebhookShow GetWebhookById(ctx, id)

A webhook endpoint

This endpoint allows the client to retrieve a single webhook given its `id`. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| Id of the webhook | 

### Return type

[**WebhookShow**](WebhookShow.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhooks

> []WebhookIndex GetWebhooks(ctx, )

Webhooks

This endpoint allows the client to retrieve a list of webhooks they have previously registered.  \"Webhooks\" are used to register HTTP endpoints that will be called once a relevant event is triggered inside the web application, events like `article_created`, `article_updated`.  It will return all webhooks, without pagination. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]WebhookIndex**](WebhookIndex.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

