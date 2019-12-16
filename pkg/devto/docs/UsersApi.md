# \UsersApi

All URIs are relative to *https://dev.to/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserAllArticles**](UsersApi.md#GetUserAllArticles) | **Get** /articles/me/all | User&#39;s all articles
[**GetUserArticles**](UsersApi.md#GetUserArticles) | **Get** /articles/me | User&#39;s articles
[**GetUserPublishedArticles**](UsersApi.md#GetUserPublishedArticles) | **Get** /articles/me/published | User&#39;s published articles
[**GetUserUnpublishedArticles**](UsersApi.md#GetUserUnpublishedArticles) | **Get** /articles/me/unpublished | User&#39;s unpublished articles



## GetUserAllArticles

> []ArticleMe GetUserAllArticles(ctx, optional)

User's all articles

This endpoint allows the client to retrieve a list of all articles on behalf of an authenticated user.  \"Articles\" are all the posts that users create on DEV that typically show up in the feed. They can be a blog post, a discussion question, a help thread etc. but is referred to as article within the code.  It will return both published and unpublished articles with pagination.  Unpublished articles will be at the top of the list in reverse chronological creation order. Published articles will follow in reverse chronological publication order.  By default a page will contain `30` articles. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetUserAllArticlesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUserAllArticlesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Pagination page. | 
 **perPage** | **optional.Int32**| Page size (defaults to 30 with a maximum of 1000). | 

### Return type

[**[]ArticleMe**](ArticleMe.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserArticles

> []ArticleMe GetUserArticles(ctx, optional)

User's articles

This endpoint allows the client to retrieve a list of published articles on behalf of an authenticated user.  \"Articles\" are all the posts that users create on DEV that typically show up in the feed. They can be a blog post, a discussion question, a help thread etc. but is referred to as article within the code.  Published articles will be in reverse chronological publication order.  It will return published articles with pagination. By default a page will contain `30` articles. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetUserArticlesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUserArticlesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Pagination page. | 
 **perPage** | **optional.Int32**| Page size (defaults to 30 with a maximum of 1000). | 

### Return type

[**[]ArticleMe**](ArticleMe.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserPublishedArticles

> []ArticleMe GetUserPublishedArticles(ctx, optional)

User's published articles

This endpoint allows the client to retrieve a list of published articles on behalf of an authenticated user.  \"Articles\" are all the posts that users create on DEV that typically show up in the feed. They can be a blog post, a discussion question, a help thread etc. but is referred to as article within the code.  Published articles will be in reverse chronological publication order.  It will return published articles with pagination. By default a page will contain `30` articles. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetUserPublishedArticlesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUserPublishedArticlesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Pagination page. | 
 **perPage** | **optional.Int32**| Page size (defaults to 30 with a maximum of 1000). | 

### Return type

[**[]ArticleMe**](ArticleMe.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserUnpublishedArticles

> []ArticleMe GetUserUnpublishedArticles(ctx, optional)

User's unpublished articles

This endpoint allows the client to retrieve a list of unpublished articles on behalf of an authenticated user.  \"Articles\" are all the posts that users create on DEV that typically show up in the feed. They can be a blog post, a discussion question, a help thread etc. but is referred to as article within the code.  Unpublished articles will be in reverse chronological creation order.  It will return unpublished articles with pagination. By default a page will contain `30` articles. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetUserUnpublishedArticlesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUserUnpublishedArticlesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Pagination page. | 
 **perPage** | **optional.Int32**| Page size (defaults to 30 with a maximum of 1000). | 

### Return type

[**[]ArticleMe**](ArticleMe.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

