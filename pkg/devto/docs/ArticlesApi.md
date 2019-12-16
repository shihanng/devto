# \ArticlesApi

All URIs are relative to *https://dev.to/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateArticle**](ArticlesApi.md#CreateArticle) | **Post** /articles | Create a new article
[**GetArticleById**](ArticlesApi.md#GetArticleById) | **Get** /articles/{id} | A published article
[**GetArticles**](ArticlesApi.md#GetArticles) | **Get** /articles | Published articles
[**GetUserAllArticles**](ArticlesApi.md#GetUserAllArticles) | **Get** /articles/me/all | User&#39;s all articles
[**GetUserArticles**](ArticlesApi.md#GetUserArticles) | **Get** /articles/me | User&#39;s articles
[**GetUserPublishedArticles**](ArticlesApi.md#GetUserPublishedArticles) | **Get** /articles/me/published | User&#39;s published articles
[**GetUserUnpublishedArticles**](ArticlesApi.md#GetUserUnpublishedArticles) | **Get** /articles/me/unpublished | User&#39;s unpublished articles
[**UpdateArticle**](ArticlesApi.md#UpdateArticle) | **Put** /articles/{id} | Update an article



## CreateArticle

> ArticleShow CreateArticle(ctx, optional)

Create a new article

This endpoint allows the client to create a new article.  \"Articles\" are all the posts that users create on DEV that typically show up in the feed. They can be a blog post, a discussion question, a help thread etc. but is referred to as article within the code.  ### Rate limiting  There is a limit of 10 articles created each 30 seconds by the same user.  ### Additional resources  - [Rails tests for Articles API](https://github.com/thepracticaldev/dev.to/blob/master/spec/requests/api/v0/articles_spec.rb) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateArticleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateArticleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **articleCreate** | [**optional.Interface of ArticleCreate**](ArticleCreate.md)| Article to create | 

### Return type

[**ArticleShow**](ArticleShow.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArticleById

> ArticleShow GetArticleById(ctx, id)

A published article

This endpoint allows the client to retrieve a single published article given its `id`.  Responses are cached for 5 minutes. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| Id of the article | 

### Return type

[**ArticleShow**](ArticleShow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArticles

> []ArticleIndex GetArticles(ctx, optional)

Published articles

This endpoint allows the client to retrieve a list of articles.  \"Articles\" are all the posts that users create on DEV that typically show up in the feed. They can be a blog post, a discussion question, a help thread etc. but is referred to as article within the code.  By default it will return featured, published articles ordered by descending popularity.  Each page will contain `30` articles.  Responses, according to the combination of params, are cached for 24 hours. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetArticlesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetArticlesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Pagination page.  This param can be used in conjuction with all other params (except when asking for fresh and rising articles by themselves).  | 
 **tag** | **optional.String**| Adding this parameter will return articles that contain the requested tag.  This param can be used by itself, with &#x60;page&#x60; or with &#x60;top&#x60;.  | 
 **username** | **optional.String**| Adding this parameter will return articles belonging to a User or Organization ordered by descending &#x60;published_at&#x60;.  If &#x60;state&#x3D;all&#x60; the number of items returned will be &#x60;1000&#x60; instead of the default &#x60;30&#x60;.  This param can be used by itself or only with &#x60;page&#x60; and &#x60;state&#x60;.  | 
 **state** | **optional.String**| Adding this will allow the client to check which articles are fresh or rising.  If &#x60;state&#x3D;fresh&#x60; the server will return published fresh articles. If &#x60;state&#x3D;rising&#x60; the server will return published rising articles.  This param can only be used by itself or with &#x60;username&#x60; if set to &#x60;all&#x60;.  | 
 **top** | **optional.Int32**| Adding this will allow the client to return the most popular articles in the last &#x60;N&#x60; days.  &#x60;top&#x60; indicates the number of days since publication of the articles returned.  This param can only be used by itself or with &#x60;tag&#x60;.  | 

### Return type

[**[]ArticleIndex**](ArticleIndex.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## UpdateArticle

> ArticleShow UpdateArticle(ctx, id, optional)

Update an article

This endpoint allows the client to updated an existing article.  \"Articles\" are all the posts that users create on DEV that typically show up in the feed. They can be a blog post, a discussion question, a help thread etc. but is referred to as article within the code.  ### Rate limiting  There are no limits on the amount of updates.  ### Additional resources  - [Rails tests for Articles API](https://github.com/thepracticaldev/dev.to/blob/master/spec/requests/api/v0/articles_spec.rb) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| Id of the article | 
 **optional** | ***UpdateArticleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateArticleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **articleUpdate** | [**optional.Interface of ArticleUpdate**](ArticleUpdate.md)| Article params to update.  *Note: if the article contains a front matter in its body, its front matter properties will still take precedence over any JSON equivalent params, which means that the full body_markdown with the modified front matter params needs to be provided for an update to be successful*  | 

### Return type

[**ArticleShow**](ArticleShow.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

