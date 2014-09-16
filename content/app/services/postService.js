'use strict';


ngBlogApp.service('postService', function ($q, $log, $http) {

    var postList = [];

    this.getPostfromPostList = function (id) {
        return currentPostSet[id];
    };
    this.updatePost = function(form) {
        var deferred = $q.defer();
        $http({ method: 'PUT', url: '/api/posts/' + form.Id , data: form}).
            success(function (data, status, headers, config) {
                deferred.resolve(status);
            }).
            error(function (data, status, headers, config) {
                deferred.reject(status);
            });
        return deferred.promise;
    };
    this.deletePost = function (id) {
        var deferred = $q.defer();
        $http({ method: 'DELETE', url: '/api/posts/' + id }).
            success(function (data, status, headers, config) {
                deferred.resolve(status);
            }).
            error(function (data, status, headers, config) {
                deferred.reject(status);
            });
        return deferred.promise;
    };
    this.createNew = function (current) {
        var deferred = $q.defer();
        $http({ method: 'POST', url: '/api/posts', data: current }).
            success(function (data, status, headers, config) {
                deferred.resolve(status);
            }).
            error(function (data, status, headers, config) {
                deferred.reject(status);
            });
        return deferred.promise;
    };
    this.fetchPost = function (post_id) {
        var deferred = $q.defer();
        $http({ method: 'GET', url: '/api/posts/' + post_id }).
            success(function (data, status, headers, config) {
                deferred.resolve(data);
            }).
            error(function (data, status, headers, config) {
                deferred.reject(data);
            });
        return deferred.promise;
    };
    this.fetchAll = function (page_num) {
        var deferred = $q.defer();
        $http({ method: 'GET', url: '/api/posts/', headers: { 'page': page_num } }).
            success(function (data, status, headers, config) {
                deferred.resolve(data);
            }).
            error(function (data, status, headers, config) {
                deferred.reject(data);
            });
        return deferred.promise;
    };

});