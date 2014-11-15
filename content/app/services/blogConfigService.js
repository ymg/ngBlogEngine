'use strict';


ngBlogApp.service('blogConfigService', function ($log, $http, $q) {
    this.fetchConfig = function () {

        var deferred = $q.defer();

        $http({method: 'GET', url: '/api/config'}).
            success(function (data, status, headers, config) {
                deferred.resolve(data);
            }).
            error(function (data, status, headers, config) {
                deferred.reject(status);
            });

        return deferred.promise;
    };
    this.fetchDbConfig = function () {
        var deferred = $q.defer();
        $http({method: 'GET', url: '/api/dbconfig'}).
            success(function (data, status, headers, config) {
                deferred.resolve(data);
            }).
            error(function (data, status, headers, config) {
                deferred.reject(status);
            });
        return deferred.promise;
    };
    this.updateDbConfig = function () {
        var deferred = $q.defer();
        $http({method: 'PUT', url: '/api/dbconfig'}).
            success(function (data, status, headers, config) {
                deferred.resolve(data);
            }).
            error(function (data, status, headers, config) {
                deferred.reject(status);
            });
        return deferred.promise;
    };
});


