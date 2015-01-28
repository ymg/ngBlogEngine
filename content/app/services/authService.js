'use strict';


ngBlogApp.service('authService', function ($rootScope, $log, $http, $q) {

    var userIsAuthenticated;

    this.auth = function (up, cb) {
        $http({method: 'POST', url: '/api/login', data: up}).
            success(function (data, status, headers, config) {
                if (status === 200) {
                    userIsAuthenticated = true;
                }
                cb(status)
            }).
            error(function (data, status, headers, config) {
                userIsAuthenticated = false;
                cb(status)
            });
    };
    this.authCheck = function () {
        var deferred = $q.defer();
        $http({method: 'PUT', url: '/api/login'}).
            success(function (data, status, headers, config) {
                deferred.resolve(status);
            }).
            error(function (data, status, headers, config) {
                deferred.reject(status);
            });
        return deferred.promise;
    };
    this.isLoggedIn = function () {
        return userIsAuthenticated;
    };
    this.logout = function (cb) {
        $http({method: 'DELETE', url: '/api/logout'}).
            success(function (data, status, headers, config) {
                userIsAuthenticated = false;
                cb(status)
            }).
            error(function (data, status, headers, config) {
                userIsAuthenticated = false;
                cb(status)
            });
    };
    this.updateCredentials = function(dat) {
        var deferred = $q.defer();
        $http({method: 'PUT', url: '/api/user/update', data: dat}).
            success(function (data, status, headers, config) {
                deferred.resolve(status);
            }).
            error(function (data, status, headers, config) {
                deferred.reject(status);
            });
        return deferred.promise;
   };
});